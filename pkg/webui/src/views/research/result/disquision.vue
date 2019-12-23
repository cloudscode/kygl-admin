<template>
  <div class="app-container">
    <div class="filter-container">
      <!--<el-input-->
      <!--:placeholder="$t('role.search')"-->
      <!--v-model="listQuery.title"-->
      <!--style="width: 200px;"-->
      <!--class="filter-item"-->
      <!--@keyup.enter.native="handleFilter"/>-->
      <el-select
        v-model="search_domain_id"
        class="filter-item"
        placeholder="Please select"
        @change="toSearch"
      >
        <el-option key="" label="全部" value="" />
        <el-option
          v-for="item in domainlist"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="toSearch"
      >
        {{ $t("table.search") }}
      </el-button>
      <el-button
        v-permission="['/disquisions:add']"
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >{{ $t("table.add") }}
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      @sort-change="sortChange"
    >
      <el-table-column
        :label="$t('disquision.id')"
        prop="id"
        align="center"
        width="65"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.title')" min-width="150px">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.affiliation_unit')">
        <template slot-scope="scope">
          {{ scope.row.affiliation_unit }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.author')">
        <template slot-scope="scope">
          {{ scope.row.authorname }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.publishing_time')">
        <template slot-scope="scope">
          {{ scope.row.publishingdate | parseTime("{y}-{m}-{d} {h}:{i}") }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.public_name')">
        <template slot-scope="scope">
          {{ scope.row.publications }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.thesis_type')">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.publication_level')">
        <template slot-scope="scope">
          {{ scope.row.levels }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('disquision.approval_status')">
        <template slot-scope="scope">
          {{ scope.row.flagapprove }}
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('role.actions')"
        align="center"
        width="200"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button
            v-permission="['/disquisions:edit']"
            type="primary"
            size="mini"
            @click="handleUpdate(scope.row)"
          >{{ $t("table.edit") }}
          </el-button>
          <el-button
            v-permission="['/disquisions:del']"
            type="danger"
            size="mini"
            @click="handleDelete(scope.row)"
          >
            {{ $t("table.delete") }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"
    />
    <DisquisionEdit
      :dialog-status="dialogStatus"
      :dialog-form-visible="dialogFormVisible"
      :disquision-data="disquisionData"
      :thesis-types="thesisTypes"
      @getList="getList"
      @changeDialogFormVisible="changeDialogFormVisible"
    ></DisquisionEdit>
  </div>
</template>

<style lang="scss" scope>
.customDialog {
  width: 1000px;
}
</style>
<script>
import {
  fetchDisquisionList,
  deleteDisquision
} from '@/api/research/result/disquision'
import {
  queryDictionaryItemsByCode
} from '@/api/dictionary/items'
import { fetchDomainList } from '@/api/domain'
import { fetchMenuList } from '@/api/menu'

import Pagination from '@/components/Pagination'
import { dataPermList } from '@/api/dataPerm'
import DisquisionEdit from '@/views/research/result/disquision_edit'

export default {
  name: `Role`,
  components: { Pagination, DisquisionEdit },
  data() {
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        skip: 0,
        limit: 20,
        q: ``
      },
      sortOptions: [
        { label: `ID Ascending`, key: `+id` },
        { label: `ID Descending`, key: `-id` }
      ],
      department: [`admin`, `editor`],
      showReviewer: false,
      disquisionData: {
        id: undefined,
        title: ``,
        orgid: ``,
        authorname: ``,
        publishingdate: ``,
        publications: ``,
        kind: `10`,
        levels: ``,
        flagapprove: ``,
        period: ``,
        volumeno: ``,
        wordlength: ``,
        knowledgeclass: ``,
        firstknowledge: ``,
        source: ``,
        publishingrange: ``,
        issn: ``,
        cn: ``
      },
      dialogFormVisible: false,
      dialogStatus: `create`,
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      downloadLoading: false,
      tree_data: [],
      tree_data_perm: [],
      tree_props: {
        children: 'children',
        label: 'name'
      },
      domainlist: [],
      domain_id: '',
      search_domain_id: '',
      menuslist: [],
      scopeType: 1,
      thesisTypes: []
    }
  },
  created() {
    this.getDeptList()
    this.getList()
    this.dictionaryItemsByCode(`DisquisionKind`)
  },
  methods: {
    dictionaryItemsByCode(code) {
      queryDictionaryItemsByCode(code).then(response => {
        const list = response.data.result
        console.log(list)
        this.thesisTypes = list
        this.disquisionData.thesis_type = '10'
      })
    },
    getList() {
      this.listLoading = true
      this.listQuery.skip = (this.listQuery.page - 1) * this.listQuery.limit
      fetchDisquisionList(this.listQuery).then(response => {
        this.list = response.data.result
        this.total = response.data.total
        console.log(this.list)
        console.log(this.total)
        // Just to simulate the time of the request
        // setTimeout(() => {
        this.listLoading = false
        // }, 1.5 * 1000)
      })
    },
    changeDialogFormVisible(vis) {
      this.dialogFormVisible = vis
    },
    toSearch() {
      if (this.search_domain_id !== '') {
        this.listQuery.q = 'd=' + this.search_domain_id
      } else {
        delete this.listQuery.q
      }
      this.getList()
    },
    // 获取权限列表 （角色权限or数据权限）
    getPermList() {
      Promise.all([
        fetchMenuList({
          q: 'd=' + this.domain_id
        }),
        dataPermList({
          start: 0,
          limit: 20,
          q: 'd=' + this.domain_id
        })
      ]).then(response => {
        console.log(response)
        const res_menus = response[0].data.result
        const result = response[1].data.result
        if (res_menus && res_menus.length > 0) {
          this.tree_data = this.o(res_menus, 0)
        } else {
          this.tree_data = []
        }
        if (result && result.length > 0) {
          this.tree_data_perm = this.o(result, 0)
        } else {
          this.tree_data_perm = []
        }
        const { menu_ids_ele = '', data_perm_ids = '' } = this.temp
        this.$refs.tree.setCheckedKeys(menu_ids_ele.split(','))
        this.$refs.treeData.setCheckedKeys(data_perm_ids.split(','))
      })
    },
    getDeptList() {
      fetchDomainList().then(response => {
        this.domainlist = response.data.result
        // Just to simulate the time of the request
        // setTimeout(() => {
        this.listLoading = false
        // }, 1.5 * 1000)
      })
    },
    o(data, id) {
      const menu = data.filter(
        o => parseInt(o.parent_id) === parseInt(id) && o.domain_id > 0
      )
      menu.forEach(o => {
        o.children = this.o(data, o.id)
      })
      return menu
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    handleModifyStatus(row, status) {
      this.$message({
        message: '操作成功',
        type: 'success'
      })
      row.status = status
    },
    sortChange(data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID(order) {
      if (order === 'ascending') {
        this.listQuery.sort = `+id`
      } else {
        this.listQuery.sort = `-id`
      }
      this.handleFilter()
    },
    resetTemp() {
      this.disquisionData = {
        id: undefined,
        title: ``,
        orgid: ``,
        authorname: ``,
        publishingdate: new Date(),
        publications: ``,
        kind: `10`,
        levels: ``,
        flagapprove: ``,
        period: ``,
        volumeno: ``,
        wordlength: ``,
        knowledgeclass: ``,
        firstknowledge: ``,
        source: ``,
        publishingrange: ``,
        issn: ``,
        cn: ``
      }
      this.domain_id = ``
    },
    handleCreate() {
      this.resetTemp()
      // this.domainlist.length > 0
      //   ? (this.domain_id = this.domainlist[0].id)
      //   : ``
      // this.getPermList()
      this.dialogStatus = `create`
      this.dialogFormVisible = true
      // this.$nextTick(() => {
      //   this.$refs['dataForm'].clearValidate()
      // })
    },

    handleUpdate(row) {
      // debugger
      // console.log('****' + JSON.stringify(row) + '***')
      this.disquisionData = Object.assign({}, row) // copy obj
      // this.temp.timestamp = new Date(this.temp.timestamp)
      // console.log('***' + this.disquisionData + '***')
      this.dialogStatus = `update`
      this.getPermList()
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    findParentMenus(menus, id) {
      const s = this.menuslist.find(i => i.id === id)
      if (s && s.parent_id !== '0') {
        menus.push(s.parent_id)
        this.findParentMenus(menus, s.parent_id)
      }
    },

    handleDelete(row) {
      this.dialogFormVisible = false
      this.$confirm(
        `<span>确认要删除吗？</span>`,
        `操作提示`,
        {
          confirmButtonText: `确认`,
          cancelButtonText: `取消`,
          type: `warning`,
          dangerouslyUseHTMLString: true
        }
      )
        .then(() => {
          deleteDisquision({ id: row.id })
            .then(() => {
              this.$notify({
                title: `成功`,
                message: `删除成功`,
                type: `success`,
                duration: 2000
              })
              // const index = this.list.indexOf(row)
              // this.list.splice(index, 1)
              this.getList()
            })
            .catch(res => {
              this.$message.error(res.msg)
            })
        })
        .catch(() => {})
    }
  }
}
</script>

<style>
.radio {
  margin-bottom: 10px;
}
.pagination-container {
  margin-top: 0;
}

.el-tree {
  box-shadow: 0px 0px 2px #e2e2e2;
}

.el-tree-node {
  display: flex;
}

.el-tree > .el-tree-node:not(:nth-last-child(2)) {
  border-bottom: 1px solid #e2e2e2;
}
.el-message-box__message span {
  color: red;
}
</style>
