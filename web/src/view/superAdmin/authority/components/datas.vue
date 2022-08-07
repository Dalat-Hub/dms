<template>
  <div>
    <div class="clearfix" style="margin: 18px">
      <el-button class="fl-right" size="small" type="primary" @click="authDataEnter">Sure</el-button>
      <el-button class="fl-left" size="small" type="primary" @click="all">Select All</el-button>
      <el-button class="fl-left" size="small" type="primary" @click="self">This Role</el-button>
      <el-button class="fl-left" size="small" type="primary" @click="selfAndChildren">This role and subroles</el-button>
    </div>
    <el-checkbox-group v-model="dataAuthorityId" @change="selectAuthority">
      <el-checkbox v-for="(item,key) in authoritys" :key="key" :label="item">{{ item.authorityName }}</el-checkbox>
    </el-checkbox-group>
    <warning-bar title="This function is only used to create the many2many relationship table of roles and roles. For specific use, you must combine the tables to realize the business. For details, please refer to the sample code (customer example)" />
  </div>
</template>

<script>
export default {
  name: 'Datas'
}
</script>

<script setup>
import { setDataAuthority } from '@/api/authority'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  },
  authority: {
    default: function() {
      return []
    },
    type: Array
  }
})

const authoritys = ref([])
const needConfirm = ref(false)
//   平铺角色
const roundAuthority = (authoritysData) => {
  authoritysData && authoritysData.forEach(item => {
    const obj = {}
    obj.authorityId = item.authorityId
    obj.authorityName = item.authorityName
    authoritys.value.push(obj)
    if (item.children && item.children.length) {
      roundAuthority(item.children)
    }
  })
}

const dataAuthorityId = ref([])
const init = () => {
  roundAuthority(props.authority)
  props.row.dataAuthorityId && props.row.dataAuthorityId.forEach(item => {
    const obj = authoritys.value && authoritys.value.filter(au => au.authorityId === item.authorityId) && authoritys.value.filter(au => au.authorityId === item.authorityId)[0]
    dataAuthorityId.value.push(obj)
  })
}

init()

// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  authDataEnter()
}

const emit = defineEmits(['changeRow'])
const all = () => {
  dataAuthorityId.value = [...authoritys.value]
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}
const self = () => {
  dataAuthorityId.value = authoritys.value.filter(item => item.authorityId === props.row.authorityId)
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}
const selfAndChildren = () => {
  const arrBox = []
  getChildrenId(props.row, arrBox)
  dataAuthorityId.value = authoritys.value.filter(item => arrBox.indexOf(item.authorityId) > -1)
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}
const getChildrenId = (row, arrBox) => {
  arrBox.push(row.authorityId)
  row.children && row.children.forEach(item => {
    getChildrenId(item, arrBox)
  })
}
// 提交
const authDataEnter = async() => {
  const res = await setDataAuthority(props.row)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'Resource setup successfully' })
  }
}

//   选择
const selectAuthority = () => {
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}

defineExpose({
  enterAndNext,
  needConfirm
})

</script>
