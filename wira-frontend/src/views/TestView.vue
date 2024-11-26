<template>
  <el-form
    ref="formRef"
    style="max-width: 600px"
    :model="numberValidateForm"
    label-width="auto"
  >
    <el-form-item
      label="email"
      prop="email"
      :rules="[
        { required: true, message: 'email is required' },
        { type: 'email', message: 'Invalid Email' },
      ]"
    >
      <el-input
        v-model="numberValidateForm.email"
        type="text"
        autocomplete="off"
      />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm(formRef)">Submit</el-button>
      <el-button @click="resetForm(formRef)">Reset</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance } from 'element-plus'

const formRef = ref<FormInstance>()

const numberValidateForm = reactive({
  age: '',
  email: '',
})

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      console.log('submit!')
    } else {
      console.log('error submit!')
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}
</script>
