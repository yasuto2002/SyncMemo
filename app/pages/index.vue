<template>
  <div
    class="m-auto w-[50%] text-center text-[25px] h-[70vh] pt-[3vw] pb-[75vh]"
  >
    <h1
      class="mb-[9vw] text-[#686868]"
      style="text-shadow: 1px 2px 3px #808080"
    >
      メモをリアルタイムで共有することで
    </h1>
    <h1
      class="mb-[8vw] text-[#686868]"
      style="text-shadow: 1px 2px 3px #808080"
    >
      チームの共同作業をもっと有効に
    </h1>
    <div class="flex items-center justify-center mb-[4vw]">
      <img
        src="assets/images/userImg.png"
        alt=""
        class="object-cover h-8 w-8"
      />
      <p class="text-[#686868]" style="text-shadow: 1px 2px 3px #808080">
        <span>0</span>ユーザー
      </p>
    </div>
    <NuxtLink
      class="mb-[6vw] bg-[#E9E5E5] px-7 py-2 text-white rounded-[10px] hover:opacity-80"
      to="/makeBoard"
    >
      start
    </NuxtLink>
  </div>
</template>
<script setup>
import { io } from "socket.io-client"
const config = useRuntimeConfig()
const authStore = useAuthStore()
const { authState } = authStore
const { authLogin } = authStore

const regData = useCookie("regData")
const logToken = useCookie("logToken")
const loginStatus = computed(() => authState.value)
</script>
