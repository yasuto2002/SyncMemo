<template>
  <div class="m-auto w-[50%] text-center text-[25px] h-[70vh] pt-[3vw] pb-[75vh]">
    <h1
      class="mb-[9vw] text-[#686868]"
      style="text-shadow: 1px 2px 3px #808080"
    >
      メモの共有化で
    </h1>
    <h1
      class="mb-[8vw] text-[#686868]"
      style="text-shadow: 1px 2px 3px #808080"
    >
      会議をもっと有益に行う
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
    <button
      class="
        mb-[6vw]
        bg-[#E9E5E5]
        px-7
        py-2
        text-white
        rounded-[10px]
        hover:opacity-80
      "
    >
      start
    </button>
  </div>
</template>
<script setup>
import { io } from "socket.io-client";
const config = useRuntimeConfig();
const socket = io(config.apiServer);
const authStore = useAuthStore()
const {authState} = authStore
const {authLogin} = authStore
// socket.emit("hello", "world");
// socket.on("news", (data) => {
//   console.log(data);
// });
// const { data } = await useFetch(`${config.apiContainer}/api`);
// let aData = data.value[0].name;
// console.log(aData);
const regData = useCookie('regData')
const logToken = useCookie('logToken')
    console.log(authState.value)
    // console.log(logToken.value)
    //     if(logToken.value){
    //     let {data}= await useFetch(`${config.apiContainer}/login`,
    //     { method: 'POST', body: { 
    //         token:logToken
    //     }}
    //     )
    //     console.log(data)
    //     }
    const loginStatus = computed(() => authState.value)
</script>
