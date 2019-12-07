<template
><el-dialog
  :title="$t('user.' + textMap[dialogStatus])"
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
    <el-form-item :label="$t('disquision.title')" prop="name">
      <el-input v-model="temp.title" />
    </el-form-item>
    <el-form-item :label="$t('role.remark')">
      <el-input v-model="temp.remark" />
    </el-form-item>
    <el-form-item :label="$t('role.domain')">
      <el-select
        v-model="domain_id"
        :disabled="dialogStatus !== 'create'"
        class="filter-item"
        placeholder="Please select"
        @change="getPermList"
      >
        <el-option
          v-for="item in domainlist"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('role.menuscope')">
      <el-radio-group v-model="scopeType" class="radio">
        <el-radio-button label="1">{{
          $t("role.menuscope")
        }}</el-radio-button>
        <el-radio-button label="2">{{
          $t("role.datascope")
        }}</el-radio-button>
      </el-radio-group>
      <el-tree
        v-show="parseInt(scopeType) === 1"
        ref="tree"
        :data="tree_data"
        :props="tree_props"
        show-checkbox
        node-key="id"
      />
      <el-tree
        v-show="parseInt(scopeType) === 2"
        ref="treeData"
        :data="tree_data_perm"
        :props="tree_props"
        show-checkbox
        node-key="id"
      />
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
