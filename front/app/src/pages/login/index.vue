<template>
  <form @submit.prevent="submit">
    <v-text-field
      v-model="name.value.value"
      :counter="10"
      :error-messages="name.errorMessage.value"
      label="用户名"
    />

    <v-text-field
      v-model="password.value.value"
      :counter="7"
      :error-messages="password.errorMessage.value"
      label="密码"
    />
    <v-text-field
      v-if="isRegister"
      v-model="rePassword.value.value"
      :counter="7"
      :error-messages="rePassword.errorMessage.value"
      label="再次输入密码"
    />

    <v-text-field
      v-if="isRegister"
      v-model="email.value.value"
      :error-messages="email.errorMessage.value"
      label="E-mail"
    />

    <v-btn
      class="me-4"
      type="submit"
    >
      {{ isRegister ? '注册' : '登录' }}
    </v-btn>

    <v-btn @click="handleReset">
      重置
    </v-btn>
  </form>
</template>

<script lang="ts" setup>
  import { useField, useForm } from 'vee-validate'
  import { R } from 'vue-router/dist/router-CWoNjPRp.mjs'
  import { LoginApi, RegisterApi } from '@/http/api/login'
  import { setToken } from '@/utils/token'
  const emits = defineEmits(['close-dialog'])
  const { isRegister } = defineProps({
    isRegister: Boolean,
  })
  const { handleSubmit, handleReset } = isRegister
    ? useForm({
      validationSchema: {
        name (value) {
          if (value?.length >= 2) return true
          return '用户名不少于两个字符'
        },
        password (value) {
          if (/^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d~!@#$%^&*()_+=\-{}[\]|:;"'<>,.?/]{7,20}$/.test(value)) return true
          return '请输入正确格式的密码'
        },
        rePassword (value) {
          if (value === password.value.value) return true
          return '两次密码不一致'
        },
        email (value) {
          if ((/^[a-z.-]+@[a-z.-]+\.[a-z]+$/i.test(value))) return true
          return '邮箱格式不正确'
        },
      },
    })
    : useForm({
      validationSchema: {
        name (value) {
          if (value?.length >= 2) return true
          return '用户名不少于两个字符'
        },
        password (value) {
          if (/^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d~!@#$%^&*()_+=\-{}[\]|:;"'<>,.?/]{7,20}$/.test(value)) return true
          return '请输入正确格式的密码'
        },
      },
    })
  const name = useField('name', undefined, {
    initialValue: '',
  })
  const password = useField('password', undefined, {
    initialValue: '',
  })
  const rePassword = useField('rePassword', undefined, {
    initialValue: '',
  })
  const email = useField('email', undefined, {
    initialValue: '',
  })

  const submit = handleSubmit(values => {
    if (isRegister) {
      RegisterApi({
        user_name: values.name,
        password: values.password,
        email: values.email,
      }).then(res => {
        setToken(res.token)
        emits('close-dialog', { msg: false })
      })
    } else {
      LoginApi({
        user_name: values.name,
        password: values.password,
      }).then(res => {
        setToken(res.token)
        emits('close-dialog', { msg: false })
      })
    }
  })
</script>
