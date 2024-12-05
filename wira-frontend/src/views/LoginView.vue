<template>
  <div class=" w-screen  mx-5">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 pl-2 text-left text-5xl tracking-tight font-medium text-gray-800 dark:text-white ">
        Welcome Back
      </h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <el-form
          ref="loginRef"
          :model="loginForm"
          :rules="loginRules"
        >
          <el-form-item label="Email" prop="email" class="block text-base font-normal" label-position="top">
            <el-input type="email" id="email" v-model="loginForm.email" />
          </el-form-item>

          <el-form-item label="Password" prop="password" class="block text-base font-normal" label-position="top">
            <el-input type="password" id="password" v-model="loginForm.password" show-password />
          </el-form-item>

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <Switch
                v-model="rememberMe"
                :class="[
                  rememberMe ? 'bg-indigo-600' : 'bg-gray-200',
                  'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2'
                ]"
              >
                <span
                  :class="[
                    rememberMe ? 'translate-x-5' : 'translate-x-0',
                    'pointer-events-none relative inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out'
                  ]"
                />
              </Switch>
              <label class="ml-2 block text-sm text-gray-900">Remember me</label>
            </div>

            <div class="text-sm">
              <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
                Forgot your password?
              </a>
            </div>
          </div>

          <div>
            <el-button class="w-full" type="primary" @click="submitForm(loginRef)" :loading="isLoading" :disabled="formBlock" >Sign In</el-button>
          </div>
        </el-form>

        <div class="mt-6">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300" />
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="bg-white px-2 text-gray-500">Or continue with</span>
            </div>
          </div>

          <div class="mt-6 grid grid-cols-2 gap-3">
            <div>
              <a
                href="#"
                class="inline-flex w-full justify-center rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-500 shadow-sm hover:bg-gray-50"
              >
                <span class="sr-only">Sign in with Google</span>
                Google
              </a>
            </div>
            <div>
              <a
                href="#"
                class="inline-flex w-full justify-center rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-500 shadow-sm hover:bg-gray-50"
              >
                <span class="sr-only">Sign in with GitHub</span>
                GitHub
              </a>
            </div>
          </div>
        </div>
        </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, computed } from 'vue'
import { Switch } from '@headlessui/vue'
import type { FormInstance, FormRules } from 'element-plus'

const loginRef = ref<FormInstance>()

interface loginInterface {
  email     : string
  password  : string
}

const rememberMe  = ref(false)
const isLoading   = ref(false)

const loginForm = reactive<FormRules<loginInterface>>({
  email     : '',
  password  : ''
})

const loginRules = reactive<FormRules<loginInterface>>({
  email: [
    {
      required: true,
      message: 'Email Required',
      trigger: 'blur',
    },
    {
      type: 'email',
      message: 'Invalid Email',
      trigger: ['blur', 'change'],
    },
  ],
  password: [
    {
      required: true,
      message: 'Password Required',
      trigger: 'blur',
    },
    {
      min: 8,
      message: 'Minimum 8 characters long',
      trigger: ['blur', 'change'],
    },
  ],
})

const submitForm = (formEl: FormInstance | undefined) => {

  console.log(formEl?.validate);

  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      console.log('submit!')
    } else {
      console.log('error submit!')
    }
  })
}

const validateEmail = () => {
  const emailRegex  = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

  if (!loginForm.email) {
    errors.value.email = 'Email is required'
  } else if (!emailRegex.test(loginForm.email)) {
    errors.value.email = 'Please enter a valid email'
  } else {
    errors.value.email = ''
  }
}

const validatePassword = () => {
  if (!loginForm.password) {
    errors.value.password = 'Password is required'
  } else if (loginForm.password.length < 6) {
    errors.value.password = 'Password must be at least 6 characters'
  } else {
    errors.value.password = ''
  }
}

const formBlock = !computed(() => {  // False: Not Email + Not PS + Have email valid + Have ps valid
  return loginForm.email
    && loginForm.password
    && !errors.value.email
    && !errors.value.password
})

const onSubmit = () => {

}

const handleSubmit = () => {
  validateEmail()
  validatePassword()

  if (formBlock.value) {
    console.log('Form submitted:', {
      rememberMe: rememberMe.value
    })
  }
}
</script>


<style scoped>
</style>
