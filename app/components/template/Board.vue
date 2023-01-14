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
<script setup lang="ts">
import { ActionCede } from "../../repository/actionCode"
const route = useRoute()
const id = route.query.id;
const router = useRouter();
const memoel = ref([]);
let memos = ref([])
const config = useRuntimeConfig()
const { $createRoom } = useNuxtApp()
$createRoom(id as string)

const ws = new WebSocket(config.socket + `/chatroom/connect?name=${id}&chatroom_id=${id}`)
ws.onopen = function () {
    console.log("接続が開かれたときに呼び出されるイベント")
}
ws.onmessage = function (event) {
  let info = JSON.parse(event.data)
  info = JSON.parse(info.data)
  if(info.actionId == ActionCede.ADD){
    let data = {
      id: info.id,
      text: info.text,
      x:info.x,
      y:info.y,
      boardId : id
    }
    memos.value.push(data)
    return
  }else{
    for(let i = 0; i < memos.value.length;i++){
      if(memos.value[i].id == info.id){
        memos.value[i].x = info.x
        memos.value[i].y = info.y
        memos.value[i].text = info.text
      }
    }
  }
}
const nuxtApp = useNuxtApp();
let apiData = await nuxtApp.$reqApi();
const coll = () => {
  let data = {
      id: "",
      text: "",
      x:0,
      y:0,
      actionId: 2,
      boardId : id
  }
  ws.send(JSON.stringify(data))
};
const moveMemo = (data) => {
  for(let i = 0; i < memos.value.length;i++){
      if(memos.value[i].id == data.id){
        memos.value[i].x = data.x
        memos.value[i].y = data.y
        memos.value[i].text = data.text
      }
    }
  ws.send(JSON.stringify(data))
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
