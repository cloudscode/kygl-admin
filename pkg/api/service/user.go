package service

import (
	"errors"
	"strconv"
	"strings"
	"zeus/pkg/api/dao"
	"zeus/pkg/api/domain/account"
	"zeus/pkg/api/domain/account/ldap"
	"zeus/pkg/api/domain/account/login"
	"zeus/pkg/api/domain/perm"
	"zeus/pkg/api/domain/user"
	"zeus/pkg/api/dto"
	"zeus/pkg/api/log"
	"zeus/pkg/api/model"
)

const pwHashBytes = 64

var userDao = dao.User{}
var userOauthDao = dao.UserOAuthDao{}
var logDao = dao.LoginLogDao{}

type UserService struct {
	//oauthdao *dao.UserOAuthDao
}

func (us UserService) InfoOfId(dto dto.GeneralGetDto) model.User {
	return userDao.Get(dto.Id, true)
}

// List - users list with pagination
func (us UserService) List(dto dto.GeneralListDto) ([]model.User, int64) {
	return userDao.List(dto)
}

// Create - create a new account
func (us UserService) Create(userDto dto.UserCreateDto) (*model.User, error) {

	//if username is exits,it can't create this user
	u := userDao.GetByUserName(userDto.Username)
	if u.Username == userDto.Username {
		return nil, errors.New("username is exits")
	}

	salt, _ := account.MakeSalt()
	pwd, _ := account.HashPassword(userDto.Password, salt)
	userModel := model.User{
		Username:     userDto.Username,
		Mobile:       userDto.Mobile,
		Password:     pwd,
		DepartmentId: userDto.DepartmentId,
		Salt:         salt,
		Sex:          userDto.Sex,
		Email:        userDto.Email,
		Title:        userDto.Title,
		Realname:     userDto.Realname,
		Status:       userDto.Status,
	}
	c := userDao.Create(&userModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
		return nil, errors.New("create user failed")
	}

	if userDto.Roles != "" {
		us.AssignRole(strconv.Itoa(userModel.Id), strings.Split(userDto.Roles, ","))
	}

	return &userModel, nil
}

// Update - update user's information
func (us UserService) Update(userDto dto.UserEditDto) int64 {
	userModel := model.User{
		Id: userDto.Id,
	}
	c := userDao.Update(&userModel, map[string]interface{}{
		"mobile":        userDto.Mobile,
		"department_id": userDto.DepartmentId,
		"status":        userDto.Status,
		"title":         userDto.Title,
		"real_name":     userDto.Realname,
		"email":         userDto.Email,
	})
	us.AssignRole(strconv.Itoa(userDto.Id), strings.Split(userDto.Roles, ","))
	return c.RowsAffected
}

// UpdateStatus - update user's status only
func (UserService) UpdateStatus(dto dto.UserEditStatusDto) int64 {
	u := userDao.Get(dto.Id, false)
	//u.Status = dto.Status
	c := userDao.Update(&u, map[string]interface{}{
		"status": dto.Status,
	})
	return c.RowsAffected
}

// UpdatePassword - update password only
func (UserService) UpdatePassword(dto dto.UserEditPasswordDto) int64 {
	salt, _ := account.MakeSalt()
	pwd, _ := account.HashPassword(dto.Password, salt)
	u := userDao.Get(dto.Id, false)
	//u.Password = pwd
	//u.Salt = salt
	c := userDao.Update(&u, map[string]interface{}{
		"password": pwd,
		"salt":     salt,
	})
	return c.RowsAffected
}

// Delete - delete user
func (us UserService) Delete(dto dto.GeneralDelDto) int64 {
	userModel := model.User{
		Id: dto.Id,
	}
	c := userDao.Delete(&userModel)
	if c.RowsAffected > 0 {
		user.DeleteUser(strconv.Itoa(dto.Id))
	}
	return c.RowsAffected
}

//VerifyAndReturnUserInfo - login and return user info
func (UserService) VerifyAndReturnUserInfo(dto dto.LoginDto) (bool, model.User) {
	userModel := userDao.GetByUserName(dto.Username)
	if login.VerifyPassword(dto.Password, userModel) {
		return true, userModel
	}
	return false, model.User{}
}

//Verfy Ldap userinfo
func (UserService) VerifyAndReturnLdapUserInfo(dto dto.LoginDto) (bool, model.User) {
	ldapConn := ldap.GetLdap()
	_, err := ldapConn.Auth(dto.Username, dto.Password)
	if err != nil {
		return false, model.User{}
	}
	return true, userDao.GetByUserName(dto.Username)
}

//AssignRoleByRoleIds - assign roles to specific user
func (UserService) AssignRoleByRoleIds(userId string, roles string) {
	// update roles
	rs := roleDao.GetRolesByIds(roles)
	var groups [][]string
	for _, role := range rs {
		groups = append(groups, []string{userId, role.RoleName})
	}
	user.OverwriteRoles(userId, groups)
}

//AssignRole - assign roles to specific user
// 这个方法同时作用与用户角色，用户用户组
func (UserService) AssignRole(userId string, roleNames []string) {
	var roles [][]string
	for _, role := range roleNames {
		if userId == "" || role == "" {
			continue
		}
		roles = append(roles, []string{userId, role})
	}
	user.OverwriteRoles(userId, roles)
}

//GetRelatedDomains - get related domains
func (UserService) GetRelatedDomains(uid string) []model.Domain {
	var domains []model.Domain
	//1.get roles by user
	roles := perm.GetGroupsByUser(uid)
	//2.get domains by roles
	for _, rn := range roles {
		role := roleDao.GetByName(rn[1])
		if role.Domain.Code == "root" {
			continue
		}
		domains = append(domains, role.Domain)
	}
	return domains
}

// GetAllRoles would return all roles of a user
func (UserService) GetAllRoles(uid string) []string {
	groups := perm.GetGroupsByUser(uid)
	roles := []string{}
	for _, group := range groups {
		roles = append(roles, group[1])
	}
	return roles
}

// GetAllPermissions - get all permission by specific user
func (UserService) GetAllPermissions(uid string) []string {
	perms := []string{}
	var path = map[string]bool{}
	for _, perm := range perm.GetAllPermsByUser(uid) {
		prefix := strings.Split(perm[1], ":")
		seg := strings.Split(prefix[0], "/")
		ss := ""
		for _, s := range seg[1:] {
			ss += "/" + s
			if ok := path[ss]; !ok {
				path[ss] = true
				perms = append(perms, ss)
			}
		}
		perms = append(perms, perm[1])
	}
	return perms
}

//GetPermissionsOfDomain - Get pure permission list  in specific domain(another backend system)
func (UserService) GetPermissionsOfDomain(uid string, domain string) []string {
	perms := perm.GetAllPermsByUser(uid)
	var polices []string
	for _, p := range perms {
		if p[3] == domain {
			polices = append(polices, p[1])
		}
	}
	return polices
}

//GetMenusOfDomain - get menus in specific domain
func (UserService) GetMenusOfDomain(uid string, domain string) []model.Role {
	roles := perm.GetGroupsByUser(uid)
	var roleNames []string
	for _, r := range roles {
		roleNames = append(roleNames, r[1])
	}
	return roleDao.GetRolesByNames(roleNames)
}

// MoveToAnotherDepartment - move users to another department
func (UserService) MoveToAnotherDepartment(uids []string, target int) error {
	return userDao.UpdateDepartment(uids, target)
}

//VerifyDTAndReturnUserInfo - verify dingtalk and return user info
func (us UserService) VerifyDTAndReturnUserInfo(code string) (user model.UserOAuth, err error) {
	dtUser, err := login.GetDingTalkUserInfo(code)
	if err != nil {
		return model.UserOAuth{}, err
	}
	User, err := userOauthDao.GetUserByOpenId(dtUser.Openid, 1)
	if err == nil {
		return User, nil
	}
	return model.UserOAuth{}, err
}

func (us UserService) UnBindUserDingtalk(from int, uid int) error {
	return userOauthDao.DeleteByUseridAndFrom(from, uid)
}

func (us UserService) GetBindOauthUserInfo(uid int) (UserInfo model.UserOAuth) {
	return userOauthDao.Get(uid)
}

//GetDomainMenu -  get specific user's menus of specific domain
func (us UserService) GetDomainMenu(uid string, domain string) []model.Menu {
	roles := us.GetAllRoles(uid)
	mids := []string{}
	for _, r := range roleDao.GetRolesByNames(roles) {
		if r.Domain.Code == domain {
			mids = append(mids, r.MenuIds)
		}
	}
	return menuDao.GetMenusByIds(strings.Join(mids, ","))
}

//CheckPermission - check user's permission in specific domain with specific policy
func (us UserService) CheckPermission(uid string, domain string, policy string) bool {
	//Could it be an alias?
	domainModel := domainDao.GetByCode(domain)
	row := menuPermAliasDao.GetByAlias(policy, domainModel.Id)
	if row.Id > 0 {
		policy = row.Perms
	}
	return perm.Enforce(uid, policy, "*", domain)
}

//insert login log
func (us UserService) InsertLoginLog(loginLogDto *dto.LoginLogDto) error {
	return logDao.Create(loginLogDto)
}
