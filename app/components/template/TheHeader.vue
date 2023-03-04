<template>
  <header class="h-[150px] w-full flex items-center mb-[2%] justify-around">
    <NuxtLink to="/"
      ><img src="~/assets/images/logo.png" class="ml-[5%]"
    /></NuxtLink>
    <div class="flex items-center justify-between w-[20%]">
      <button v-if="loginStatus" @click="logout">ログアウト</button>
      <NuxtLink to="/login" v-if="!loginStatus" class="">ログイン</NuxtLink>
      <NuxtLink to="/reg" v-if="!loginStatus">登録</NuxtLink>
    </div>
  </header>
</template>
<script setup lang="ts">
import type { errCode } from "~~/repository/errCode"
const authStore = useAuthStore()
const http = useHttp()
const router = useRouter()
const { authState, authLogout } = authStore
const loginStatus = ref(computed(() => authState.value))
const { $tokenDelete } = useNuxtApp()
const logout = async () => {
  const token = useCookie<{ token: string }>("token")
  if (typeof token.value != "undefined" || token.value != null) {
    let errcode = await $tokenDelete(token.value.token)
  }
  await authLogout()
  token.value.token = null
  router.push("/")
}
</script>
