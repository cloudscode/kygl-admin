<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select
        v-model="domain_id"
        class="filter-item"
        placeholder="Please select"
        @change="getList"
      >
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
        @click="getList"
      >{{ $t("table.search") }}</el-button
      >
      <el-button
        v-permission="['/disquisions:add']"
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >{{ $t("table.add") }}</el-button
      >
    </div>
    <template v-if="domain_id && list">
      <tree-table
        :data="list"
        :eval-func="func"
        :eval-args="args"
        :columns="columns"
        :expand-all="expandAll"
        @loadChilden="loadChilden"
      >
        <el-table-column
          :label="$t('disquisions.actions')"
          width="200"
          align="center"
        >
          <template v-if="scope.row.id" slot-scope="scope">
            <el-button
              v-permission="['/disquisions:edit']"
              type="text"
              @click="handleUpdate(scope.row)"
            >
              {{ $t("table.edit") }}
            </el-button>
            <el-button
              v-permission="['/disquisions:del']"
              type="text"
              @click="handleDelete(scope.row)"
            >
              {{ $t("table.delete") }}
            </el-button>
          </template>
        </el-table-column>
      </tree-table>
    </template>
    <el-dialog
      :name="$t('dictionaries.' + textMap[dialogStatus])"
      :visible.sync="dialogFormVisible"
    >
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="80px"
        style=" margin: 0 50px;"
      >
        <el-form-item :label="$t('dictionaries.name')" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item :label="$t('dictionaries.parents')">
          <el-cascader
            :style="{ width: '100%' }"
            :options="data_options"
            :props="cascader_props"
            v-model="temp.parents"
            change-on-select
          />
        </el-form-item>
        <el-form-item :label="$t('dictionaries.code')" prop="code">
          <el-input v-model="temp.code" />
        </el-form-item>
        <el-form-item
          v-if="temp.menu_type !== '2'"
          :label="$t('dictionaries.sort')"
          prop="sortvalue"
        >
          <el-input v-model="temp.sortvalue" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{
          $t("table.cancel")
        }}</el-button>
        <el-button
          type="primary"
          @click="dialogStatus === 'create' ? createData() : updateData()"
        >{{ $t("table.confirm") }}</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
import treeTable from '@/components/AsyncTreeTable'
import treeToArray from '@/directive/customEval'
import icons from '@/utils/requireIcons'
import {
  fetchDictionaryItemsList,
  createDictionaryItems,
  updateDictionaryItems,
  deleteDictionaryItems
} from '@/api/dictionary/items'
import { fetchDomainList } from '@/api/domain'

export default {
  name: 'DictionaryItems',
  components: { treeTable },
  filters: {
    // tag(type) {
    //   return type === 0 ? '目录' : type === 1 ? '菜单' : '权限'
    // }
  },
  data() {
    return {
      func: treeToArray,
      expandAll: false,
      list: {},
      args: [null, null, 'timeLine'],
      columns: [
        {
          value: 'name',
          text: this.$t('menu.name')
        }
      ],
      //  toggleExpandedFunction: toggleCurrentExpanded,
      temp: {
        id: undefined,
        name: '',
        menu_type: '1',
        code: '',
        sortvalue: 1,
        perms: '',
        icon: '',
        parents: ['0']
      },
      icons: icons,
      dialogFormVisible: false,
      dialogStatus: 'create',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      rules: {
        type: [
          { required: true, message: 'type is required', trigger: 'change' }
        ],
        timestamp: [
          {
            type: 'date',
            required: true,
            message: 'timestamp is required',
            trigger: 'change'
          }
        ],
        name: [{ required: true, message: '名称必须填写', trigger: 'blur' }],
        // url: [{ required: true, message: '路由链接必须填写', trigger: 'blur' }],
        sortvalue: [
          { required: true, message: '排序必须填写', trigger: 'blur' }
        ]
      },
      index: 1,
      data_copy: [],
      data_options: [],
      cascader_props: {
        value: 'id',
        label: 'name'
      },
      domainlist: [],
      domain_id: ''
    }
  },

  mounted() {
    // this.list = this.recursive(asyncRouterMap, {})
    // this.recursive(asyncRouterMap, {})
    // this.os(this.list, 0)
    // this.oks(0)
  },
  created() {
    // this.getList()
    this.getDeptList()
  },
  methods: {
    iconShow(index, record) {
      console.log('parent_id' + record)
      return true
    },
    // 点击按钮后置空From
    typeChange(type) {
      this.resetTemp()
      this.temp.menu_type = type
    },
    os(data) {
      data.forEach(o => {
        this.data_copy.push(o)
        if (o.children && Array.isArray(o.children)) {
          this.os(o.children)
        }
      })
    },
    oks(i) {
      if (i < this.data_copy.length) {
        var _self = this
        this.autoCreate(this.data_copy[i]).then(res => {
          _self.oks(i + 1)
        })
      }
    },
    getDeptList() {
      this.listLoading = true
      fetchDomainList().then(res => {
        this.domainlist = res.data.result
        if (
          this.domain_id === '' &&
          this.domainlist &&
          this.domainlist.length > 0
        ) {
          this.domain_id = this.domainlist[0].id
          this.getList()
        }
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    loadChilden(params) {
      this.listLoading = true
      var row = params.record
      var parent_id = row.id
      fetchDictionaryItemsList({ q: 'd=' + parent_id }).then(res => {
        const res_menus = res.data.result

        row._expanded = true

        this.$set(row, 'children', treeToArray(res_menus, false, row, 2, null))

        // Just to simulate the time of the request
        // setTimeout(() => {
        this.listLoading = false
        // }, 1.5 * 1000)
      })
    },
    toggleExpanded(row) {
      //  console.log("trIndex"+trIndex)
      // var row = this.formatData[trIndex]
      // console.log(row)
      this.listLoading = true
      var parent_id = row.id
      console.log('parent_id' + parent_id)
      fetchDictionaryItemsList({ q: 'd=' + parent_id }).then(res => {
        const res_menus = res.data.result

        row._expanded = true

        this.$set(row, 'children', treeToArray(res_menus, false, row, 2, null))

        // Just to simulate the time of the request
        // setTimeout(() => {
        this.listLoading = false
        // }, 1.5 * 1000)
      })
    },
    getList() {
      this.listLoading = true
      if (this.parent_id === '' || this.parent_id === undefined) {
        this.parent_id = 0
      }
      fetchDictionaryItemsList({ q: 'd=' + this.parent_id }).then(res => {
        const res_menus = res.data.result
        // console.log(res_menus)
        // console.log(this.o(res_menus, 0))
        // const menu = this.o(res_menus, 0)
        if (res_menus && res_menus.length > 0) {
          this.list = this.o(res_menus, 0)
          this.data_options = this.list
        } else {
          this.list = []
          this.data_options = this.list
        }
        // Just to simulate the time of the request
        // setTimeout(() => {
        this.listLoading = false
        // }, 1.5 * 1000)
      })
    },
    o(data, id) {
      console.log(data)
      data.forEach(o => {
        o._expanded = false
        o._level = 1
        o.children = [{}]

        // const children = this.o(data, o.id)
        // if (children && children.length > 0) {
        //   o.children = children
        // }
      })
      console.log(data)
      return data
      // const menu = data.filter(o => o.parent_id === id)
      // menu.forEach(o => {
      //   o._expanded= false
      //   o._level= 1
      //   console.log(o.id)
      //   const children = this.o(data, o.id)
      //   if (children && children.length > 0) {
      //     o.children = children
      //   }
      // })
      // return menu
    },
    recursive(obj, parent) {
      const output = []
      let temp = []
      let index = 1
      obj.forEach(o => {
        if (o.permission || o.permission !== false) {
          temp = {
            name: this.$t('route.' + o.meta.title),
            label: this.$t('route.' + o.meta.title),
            icon: o.meta.icon,
            id: this.index,
            value: this.index,
            sortvalue: index,
            url:
              JSON.stringify(parent) === '{}'
                ? o.path
                : parent.url + '/' + o.path,
            parent_id: JSON.stringify(parent) === '{}' ? 0 : parent.id,
            perms: '',
            alias: '',
            menu_type: '1'
          }
          this.index += 1
          index += 1
          if (o.children && Array.isArray(o.children)) {
            temp.children = this.recursive(o.children, temp)
          }
          if (o.auth && Array.isArray(o.auth)) {
            const items = []
            let item = {}
            let item_index = 1
            o.auth.forEach(o => {
              item = {
                name: o.name,
                label: o.name,
                icon: '',
                id: this.index,
                value: this.index,
                sortvalue: item_index,
                url: '',
                parent_id: temp.id,
                perms: temp.url + ':' + o.code,
                menu_type: '2'
              }
              items.push(item)
              this.index += 1
              item_index += 1
            })
            temp.children = items
          }
          output.push(temp)
        }
      })
      return output
    },
    findParent(id, output, data) {
      data.forEach(o => {
        if (o.id === id) {
          output.push(o.parent_id)
          this.findParent(o.parent_id, output, this.list)
        }
        if (o.children && Array.isArray(o.children)) {
          this.findParent(id, output, o.children)
        }
      })
      return output
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        name: '',
        menu_type: '1',
        url: '',
        sortvalue: 1,
        perms: '',
        icon: '',
        parents: ['0']
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row) // copy obj
      console.log(this.temp)
      this.temp.created_time = new Date(this.temp.created_time)
      this.temp.last_update_time = new Date(this.temp.last_update_time)
      this.temp.parents = ['0']
      this.findParent(this.temp.id, this.temp.parents, this.list)
      // this.temp.parents.reverse()
      this.dialogStatus = `update`
      this.dialogFormVisible = true
    },
    handleDelete(row) {
      this.$confirm(`此操作将永久删除数据, 是否继续?`, `提示`, {
        confirmButtonText: `确定`,
        cancelButtonText: `取消`,
        type: `warning`
      })
        .then(() => {
          deleteDictionaryItems({ id: row.id })
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
    },
    createData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          this.temp.parent_id = this.temp.parents[this.temp.parents.length - 1]
          delete this.temp.parent
          createDictionaryItems(this.temp)
            .then(() => {
              // this.list.unshift(this.temp)
              this.getList()
              this.dialogFormVisible = false
              this.$notify({
                name: `成功`,
                message: `创建成功`,
                type: 'success',
                duration: 2000
              })
            })
            .catch(res => {
              this.$message.error(res.msg)
            })
        }
      })
    },
    autoCreate(obj) {
      this.temp = obj
      this.temp.name = obj.name
      this.temp.parent_id = obj.parent_id
      delete this.temp.children
      return createDictionaryItems(this.temp)
    },
    updateData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          delete this.temp.children
          const tempData = Object.assign({}, this.temp)
          tempData.parent_id = tempData.parents[this.temp.parents.length - 1]
          delete tempData.parent
          updateDictionaryItems(tempData)
            .then(() => {
              this.getList()
              this.dialogFormVisible = false
              this.$notify({
                name: `成功`,
                message: `更新成功`,
                type: `success`,
                duration: 2000
              })
            })
            .catch(res => {
              this.$message.error(res.msg)
            })
        }
      })
    }
  }
}
</script>
