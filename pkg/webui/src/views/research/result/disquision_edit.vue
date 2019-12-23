<template>
  <el-dialog
    :title="$t('user.' + textMap[dialogStatus])"
    :visible.sync="dialogFormVisible"
    custom-class="customDialog"
  >
    <el-form
      ref="dataForm"
      :rules="rules"
      :model="disquisionData"
      label-position="right"
      label-width="100px"
      style=" margin: 0 50px;"
    >
      <el-form-item :label="$t('disquision.thesis_type')" prop="thesis_type">
        <el-radio-group v-model="disquisionData.kind">
          <el-radio-button v-for="thesisType in thesisTypes" :label="thesisType.id" :key="thesisType.id" >{{ thesisType.name }}</el-radio-button >
        </el-radio-group>
      </el-form-item>
      <el-form-item :label="$t('disquision.title')" prop="title">
        <el-input v-model="disquisionData.title" />
      </el-form-item>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.publication_type')" prop="publication_type">
            <el-input v-model="disquisionData.period" ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.publishing_time')" class="lableLeft" prop="publishing_time">
            <!-- <el-input v-model="disquisionData.publishingdate" ></el-input> -->
            <el-date-picker
              v-model="disquisionData.publishingdate"
              style="width:100%"
              type="datetime"
              format="yyyy-MM-dd hh:mm"
              value-format="yyyy-MM-dd hh:mm"
              placeholder="选择日期时间">
            </el-date-picker>
          </el-form-item>

        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.volume_no')">
            <el-input v-model="disquisionData.volumeno" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.cn_num')">
            <el-input v-model="disquisionData.cn" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.knowledge_class')">
            <el-input v-model="disquisionData.knowledgeclass" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.first_knowledge')">
            <el-input v-model="disquisionData.firstknowledge" />
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
            <el-input v-model="disquisionData.publishingrange" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.affiliation_unit')">
            <el-input v-model="disquisionData.orgid" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.author')">
            <el-input v-model="disquisionData.authorname" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.publication_level')">
            <el-input v-model="disquisionData.levels" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.public_name')">
            <el-input v-model="disquisionData.publications" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row type="flex" class="row-bg">
        <el-col :span="12">
          <el-form-item :label="$t('disquision.word_length')">
            <el-input v-model="disquisionData.wordlength" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('disquision.issn_num')">
            <el-input v-model="disquisionData.issn" />
          </el-form-item>
        </el-col>
      </el-row>

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
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      rules: {
        // type: [{ required: true, message: 'type is required', trigger: 'change' }],
        // timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        // title: [{ required: true, message: '角色名必须填写', trigger: 'blur' }]
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
          this.disquisionData.publishingdate = +new Date(this.disquisionData.publishingdate)
          console.log('****' + JSON.stringify(this.disquisionData) + '***')
          createDisquision(this.disquisionData)
            .then(() => {
              // this.getList()
              this.$emit('getList', {})
              this.$emit('changeDialogFormVisible', false)

              this.tree_data = []

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
          const tempData = Object.assign({}, this.disquisionData)
          tempData.timestamp = +new Date(tempData.timestamp)
          // console.log('****' + JSON.stringify(tempData) + '***')
          updateDisquision(tempData)
            .then(() => {
              this.$emit('getList', {})
              this.$emit('changeDialogFormVisible', false)
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

