<template>
  <header class="h-[150px] w-full flex items-center mb-[2%] justify-around">
    <NuxtLink to="/"><img src="~/assets/images/logo.png" class="ml-[5%]" /></NuxtLink>
    <div class="flex items-center">
      <button v-if="loginStatus" @click="logout">ログアウト</button>
      <NuxtLink to="/login" v-if="!loginStatus" @click="logout">ログイン</NuxtLink>
    </div>
  </header>
</template>
<script setup lang="ts">
  const authStore = useAuthStore()
  const {authState, authLogout} = authStore
  const loginStatus = ref(computed(() => authState.value))
  const logout = async() =>{
    await  authLogout()
    const token = useCookie<{ token: string}>("token",{maxAge: 0})
    token.value = null
  }
</script>

