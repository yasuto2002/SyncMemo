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
      @moveMemo="moveMemo"
    />
  </div>
</template>
<script setup>
const route = useRoute()
const id = route.query.id;
const router = useRouter();
const memoel = ref([]);
let memos = ref([])
const createRoom = async() =>{
    const { data, pending, refresh, error }  = await useFetch(`http://localhost:8080/chatroom/create/${id}`, { method: 'POST' });
    if(error.value){
        console.log(error.value)
        router.push("/error")
        return
    }
    let re = JSON.parse(data.value)
    if(re.status != 'success'){
      return
    }
    refresh()
}
createRoom()

const config = useRuntimeConfig()
const ws = new WebSocket(config.socket + `/chatroom/connect?name=${id}&chatroom_id=${id}`)
ws.onopen = function () {
    console.log("接続が開かれたときに呼び出されるイベント")
    send()
}
ws.onmessage = function (event) {
  let info = JSON.parse(event.data)
  info = JSON.parse(info.data)
  memos.value[0].x = info.x
  memos.value[0].y = info.y
  console.log(memos.value)
}
var send =() => {
    let send_msg = "aaaa"
    ws.send(send_msg)
}

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
      x:10,
      y:10,
  })
  console.log(memos.value)
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
const moveMemo = (data) => {
  memos.value[0].x = data.x
  memos.value[0].y = data.y
  ws.send(JSON.stringify(data))
  console.log(memos.value)
};

onMounted( async() => {
})
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
