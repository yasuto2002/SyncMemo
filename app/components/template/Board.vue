<template>
  <div
    class="w-[1200px] h-[1000px] bg-[#F3F3F3] m-auto"
    style="overflow: hidden; position: relative; box-sizing: border-box"
  >
    <PartsMemo
      class="drag-and-drop bg-[#fffdfd] text-[#7c4e4e] rounded-[10px] el absolute"
      v-for="memo in memos"
      :key="memo"
      :data="{ memo }"
      :ref="
        (el) => {
          if (el) memoel[0] = el;
        }
      "
    />
  </div>
</template>
<script setup>
const route = useRoute()
const id = route.query.id;
// import { io } from "socket.io-client";
// const config = useRuntimeConfig();
// const socket = io(config.apiServer);
// const sendPosition = (position) => {
//   socket.emit("pass", position);
// }
// const config = useRuntimeConfig();
const nuxtApp = useNuxtApp();
// let { socket, makeMemo, sendMemo } = await nuxtApp.$makeSoket();
let apiData = await nuxtApp.$reqApi();
// console.log(apiData);
let memos = ref([])
// const getMemos = async() =>{
//   let {data}= await useFetch(`${config.apiServer}/getMemo`
//   )
//   if (!data.value) {
//     clearError({ redirect: '/' })
//   }else{
//     console.log(data.value.memos)
//     memos.value = data.value.memos
//   }
// }
// getMemos()
const coll = () => {
  // socket.emit("makeMemo");
  // console.log(1)
  memos.value.push({
      id: 0,
      text: "",
      x:0,
      y:0,
  })
};
// socket.on("preservation", (data) => {
//   id = data;
// });
// socket.on("addMemo", (data) => {
//   memos.value = data.memos;
//   console.log(memos.value[0]._id)
// });
// socket.on("receiveData", (receiveData) => {
//   console.log(receiveData)
//   memos.value = receiveData.memos;
//   refresh()
// });
// const moveMemo = (data) => {
//   socket.emit("moveMemo",data);
// };
const memoel = ref([]);
onMounted(() => {
  console.log(id)
});
defineExpose({
  coll,
});
// watchEffect(() => {});
// const refresh = () => refreshNuxtData('memos')
</script>
<style scoped>
.drag-and-drop {
  cursor: move;
  position: absolute;
}

.SetTextarea {
  height: 100%;
  width: 100%;
}
.el{
  width: 200px;
      height: 200px;
      white-space: pre-wrap;
      display: -webkit-box;
      -webkit-box-orient: vertical;
      -webkit-line-clamp: 2;
      overflow: hidden;
}
</style>
