<template>
  <div>
    <PartsModalUrl></PartsModalUrl>
    <div class="flex justify-between w-[40%] mb-[1%] ml-[5%]">
      <PartsBoardName>{{ route.query.name }}</PartsBoardName>
      <label for="my-modal" class="text-[#D8CFCF]">共有</label>
      <button @click="add" class="text-[#D8CFCF]">メモ追加</button>
    </div>
    <div
      class="w-[1200px] h-[1000px] bg-[#F3F3F3] m-auto"
      style="overflow: hidden; position: relative; box-sizing: border-box"
    >
      <PartsMemo
        class="drag-and-drop bg-[#fffdfd] text-[#7c4e4e] rounded-[10px] el absolute"
        v-for="(memo, i) in memos"
        :key="i"
        :data="{ memo }"
        :boardId="boardId"
        :do="ary[i]"
        :ref="
          (el) => {
            if (el) memoel[0] = el
          }
        "
        @moveMemo="moveMemo"
      />
    </div>
  </div>
</template>
<script setup lang="ts">
import { ActionCede } from "../../repository/actionCode"
import type { errCode } from "../../repository/errCode"
import type { Ref } from "vue"
import type { Memo } from "../../repository/respons/memo"
import type { SendMemo } from "../../repository/request/sendMemo"
import type { createRoom } from "~~/repository/respons/createRoom"

const route = useRoute()
const router = useRouter()
const config = useRuntimeConfig()
const { $createRoom } = useNuxtApp()
const { $memos } = useNuxtApp()
const http = useHttp()

const boardId: Ref<string> = ref(route.query.id as string)
const memoel = ref([])
let memos: Ref<Array<Memo>> = ref(new Array())
const stetusCode: Ref<errCode> = ref(200)

let [_, code] = await $createRoom(boardId.value as string)
switch (code) {
  case http.value.InternalServerError:
    router.push("/error")
    break
  case http.value.Unauthorized:
    router.push("/error")
    break
  case http.value.Unauthorized:
    router.push("/login")
    break
}
let evacuation: Array<Memo>
;[evacuation, stetusCode.value] = await $memos(boardId.value as string)

switch (code) {
  case http.value.InternalServerError:
    router.push("/error")
    break
  case http.value.BadRequest:
    router.push("/error")
    break
  case http.value.Unauthorized:
    router.push("/login")
    break
  default:
    if (stetusCode.value == 200 && evacuation != null) {
      memos.value = evacuation
    } else if (evacuation == null) {
      memos.value = new Array()
    }
}

const ary: Ref<boolean[]> = ref(
  new Array<boolean>(memos.value.length).fill(false)
)

const ws = new WebSocket(
  config.socket +
    `/chatroom/connect?name=${boardId.value}&chatroom_id=${boardId.value}`
)

ws.onopen = function () {
  console.log("接続が開かれたときに呼び出されるイベント")
}
ws.onmessage = function (event: MessageEvent) {
  let info = JSON.parse(event.data)
  info = JSON.parse(info.data)

  switch (info.actionId) {
    case ActionCede.ADD:
      let data: Memo = {
        id: info.id,
        text: info.text,
        x: info.x,
        y: info.y,
        boardid: boardId.value as string,
      }
      memos.value.push(data)
      ary.value.push(false)
      break
    case ActionCede.START:
      for (let i = 0; i < memos.value.length; i++) {
        if (memos.value[i].id == info.id) {
          cheng(i, info)
          ary.value[i] = true
        }
      }
      break
    case ActionCede.INPUT:
      for (let i = 0; i < memos.value.length; i++) {
        if (memos.value[i].id == info.id) {
          cheng(i, info)
          ary.value[i] = true
        }
      }
      break
    case ActionCede.END:
      for (let i = 0; i < memos.value.length; i++) {
        if (memos.value[i].id == info.id) {
          cheng(i, info)
          ary.value[i] = false
        }
      }
      break
    case ActionCede.DELETE:
      deletMemo(info.id)
      break
    default:
      for (let i = 0; i < memos.value.length; i++) {
        if (memos.value[i].id == info.id) {
          cheng(i, info)
        }
      }
  }
}

//受け取ったメモデータを反映
const cheng = (i: number, info) => {
  memos.value[i].x = info.x
  memos.value[i].y = info.y
  memos.value[i].text = info.text
}

//メモ削除
const deletMemo = (id: string) => {
  for (let i = 0; i < memos.value.length; i++) {
    if (memos.value[i].id == id) {
      memos.value.splice(i, 1)
    }
  }
}

const add = () => {
  let data: SendMemo = {
    id: "",
    text: "",
    x: 0,
    y: 0,
    actionId: ActionCede.ADD,
    boardId: boardId.value,
  }
  ws.send(JSON.stringify(data))
}
const moveMemo = (data: SendMemo) => {
  for (let i = 0; i < memos.value.length; i++) {
    if (memos.value[i].id == data.id) {
      memos.value[i].x = data.x
      memos.value[i].y = data.y
      memos.value[i].text = data.text
    }
  }
  if (data.actionId == ActionCede.DELETE) {
    deletMemo(data.id)
  }
  ws.send(JSON.stringify(data))
}

onMounted(async () => {})
defineExpose({
  add,
})
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
.el {
  width: 200px;
  height: 200px;
  white-space: pre-wrap;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
}
</style>
