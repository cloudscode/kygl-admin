<template>
  <el-dialog
    :title="$t('user.' + textMap[dialogStatus])"
    :visible.sync="dialogFormVisible"
    custom-class="customDialog"
  >
    <el-form
      ref="dataForm"
      :rules="rules"
      :model="temp"
      label-position="left"
      label-width="100px"
      style=" margin: 0 50px;"
    >
      <el-form-item :label="$t('disquision.thesis_type')" prop="thesis_type">
        <el-radio-group v-model="disquisionData.thesis_type">
          <el-radio-button v-for="thesisType in thesisTypes" :label="thesisType.id" :key="thesisType.id" >{{ thesisType.name }}</el-radio-button >
        </el-radio-group>
      </el-form-item>
      <el-form-item :label="$t('disquision.title')" prop="title">
        <el-input v-model="disquisionData.title" />
      </el-form-item>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.publication_type')" prop="publication_type">
            <el-input v-model="disquisionData.publication_type" ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.publishing_time')" prop="publishing_time">
            <el-input v-model="disquisionData.publishing_time" ></el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.volume_no')">
            <el-input v-model="disquisionData.volume_no" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.cn_num')">
            <el-input v-model="disquisionData.cn_num" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.knowledge_class')">
            <el-input v-model="disquisionData.knowledge_class" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.first_knowledge')">
            <el-input v-model="disquisionData.first_knowledge" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.source')">
            <el-input v-model="disquisionData.source" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.publishing_range')">
            <el-input v-model="disquisionData.publishing_range" />
          </el-form-item>
        </el-col>
      </el-row>
      <!-- <el-row type="flex" class="row-bg">
        <el-col :span="12">
        </el-col>
        <el-col :span="12">
        </el-col>
      </el-row> -->
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.affiliation_unit')">
            <el-input v-model="disquisionData.affiliation_unit" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.author')">
            <el-input v-model="disquisionData.author" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.publication_level')">
            <el-input v-model="disquisionData.publication_level" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.public_name')">
            <el-input v-model="disquisionData.public_name" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.word_length')">
            <el-input v-model="disquisionData.word_length" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.issn_num')">
            <el-input v-model="disquisionData.issn_num" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item :label="$t('disquision.publication_type')">
        <el-input v-model="disquisionData.publication_type" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="(dialogFormVisible = false), (tree_data = [])">{{
        $t("table.cancel")
      }}</el-button>
      <el-button
        type="primary"
        @click="dialogStatus === 'create' ? createData() : updateData()"
      >{{ $t("table.confirm") }}
      </el-button>
    </div>
  </el-dialog>
</template>

<script>
import {

  createDisquision,
  updateDisquision

} from '@/api/research/result/disquision'
// import treeToArray from './eval'
export default {
  name: 'DisquisionEdit',
  props: {
    dialogStatus: {
      type: String,
      default: `create`
    },
    disquisionData: {
      type: Object,
      default: null
    },
    dialogFormVisible: {
      type: Boolean,
      default: false
    },
    thesisTypes: {
      type: Array,
      default: null
    }
  },
  data() {
    return {
      // temp: {
      //   id: undefined,
      //   title: ``,
      //   remark: ``,
      //   role_name: ``,
      //   menu_ids: [],
      //   domain_id: ``,
      //   thesis_type: ``,
      //   issn_num: ``,
      //   cn_num: ``,
      //   publishing_range: ``
      // }
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      rules: {
        // type: [{ required: true, message: 'type is required', trigger: 'change' }],
        // timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: '角色名必须填写', trigger: 'blur' }]
        // remark: [{ required: true, message: '备注必须填写', trigger: 'blur' }]
      }
    }
  },
  computed: {

  },
  methods: {
    createData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          // this.disquisionData.id = parseInt(Math.random() * 100) + 1024 // mock a id
          this.disquisionData.author = `vue-element-admin`
          this.disquisionData.role_name = this.disquisionData.name
          this.disquisionData.menu_ids_ele = this.$refs.tree.getCheckedKeys().join(',')
          const menu_ids = []
          const data_perm_ids = []
          this.$refs.tree.getCheckedKeys().forEach(o => {
            menu_ids.push(o)
          })
          this.$refs.treeData.getCheckedKeys().forEach(o => {
            data_perm_ids.push(o)
          })
          this.disquisionData.menu_ids = Array.from(new Set(menu_ids)).join(',')
          this.disquisionData.data_perm_ids = Array.from(new Set(data_perm_ids)).join(
            ','
          )
          this.disquisionData.domain_id = this.domain_id
          this.disquisionData.thesis_type = `10`
          createDisquision(this.temp)
            .then(() => {
              // this.list.unshift(this.temp)
              this.getList()
              this.tree_data = []
              this.dialogFormVisible = false
              this.$notify({
                title: `成功`,
                message: `创建成功`,
                type: `success`,
                duration: 2000
              })
            })
            .catch(res => {
              this.$message.error(res.msg)
            })
        }
      })
    },
    updateData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          this.disquisionData.role_name = this.disquisionData.name
          this.disquisionData.menu_ids_ele = this.$refs.tree.getCheckedKeys().join(',')
          const menu_ids = []
          const data_perm_ids = []
          this.$refs.tree.getCheckedKeys().forEach(o => {
            menu_ids.push(o)
            this.findParentMenus(menu_ids, o)
            // const s = this.menuslist.find(i => i.id === o)
            // if (s) {
            //   menu_ids.push(s.parent_id)
            // }
          })
          this.$refs.treeData.getCheckedKeys().forEach(o => {
            data_perm_ids.push(o)
          })
          this.disquisionData.menu_ids = Array.from(new Set(menu_ids)).join(',')
          this.disquisionData.data_perm_ids = Array.from(new Set(data_perm_ids)).join(
            ','
          )
          this.disquisionData.domain_id = this.domain_id
          delete this.disquisionData.domain
          const tempData = Object.assign({}, this.temp)
          tempData.timestamp = +new Date(tempData.timestamp)
          updateDisquision(tempData)
            .then(() => {
              this.getList()
              this.tree_data = []
              this.dialogFormVisible = false
              this.$notify({
                title: `成功`,
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
