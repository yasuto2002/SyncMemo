<template>
  <header class="h-[150px] w-full flex items-center mb-[2%] justify-between">
    <NuxtLink to="/"><img src="~/assets/images/logo.png" class="ml-[5%]" /></NuxtLink>
    <div class="flex items-center">
      <button v-if="loginStatus" @click="logout">ログアウト</button>
    </div>
  </header>
</template>
<script setup>
  const authStore = useAuthStore()
  const {authState, authLogout} = authStore
  const loginStatus = ref(computed(() => authState.value))
  const logout = async() =>{
    console.log(`logout`)
    await  authLogout()
    const logToken =  useCookie('logToken',0)
    logToken.value = null
    console.log(loginStatus)
  }
</script>

