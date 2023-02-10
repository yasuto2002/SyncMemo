<template>
  <div
    class="drag-and-drop bg-[#fffdfd] text-[#7c4e4e] rounded-[10px] elv shadow-blue-300 relative"
    id="red-box"
    ref="memo"
    :class='{"shadow-2xl":opponent,"opacity-50":opponent}'
  >
    <p v-if="view" class="w-full h-full p-[10%]">{{ text }}</p>
    <textarea
      name=""
      id=""
      v-model="text"
      v-if="!view"
      class="resize-none focus:outline-none w-full h-full p-[10%]"
      @blur="out"
      @keyup="input"
    ></textarea>
    <div class="absolute bg-red-400 top-0 right-0 rounded-[10%] opacity-70 hover:cursor-pointer" :class="cloer()" @click="deleteMome">
      &times;
    </div>
  </div>
</template>
<script setup lang="ts">

import { ActionCede } from "../../repository/actionCode"
import type { SendMemo } from '../../repository/request/sendMemo'
import type { Ref } from 'vue'
const opponent:Ref<boolean> = ref(false)
const memo:Ref<HTMLElement> = ref(null) // refの値をとる
let view:Ref<boolean> = ref(true);
let text:Ref<string> = ref()
const emit = defineEmits<{
  (event: "moveMemo", firstName: object): void;
}>()
const props = defineProps({
  data: { type: Object, required: true },
  boardId:{ type: String, required: true},
  do:{ type: Boolean, required: true},
});

text.value = props.data.memo.text

const out = () => {
  view.value = !view.value;
  let memoData:SendMemo = makeSendMemo(ActionCede.END)
  emit("moveMemo",memoData)
};
const input = () =>{
  let memoData:SendMemo = makeSendMemo(ActionCede.INPUT)
  emit("moveMemo",memoData)
}

const move = () =>{
  let memoData:SendMemo = makeSendMemo(ActionCede.MOVE)
  emit("moveMemo",memoData)
}

const start = () =>{
  let memoData:SendMemo = makeSendMemo(ActionCede.START)
  emit("moveMemo",memoData)
}

const end = () =>{
  let memoData:SendMemo = makeSendMemo(ActionCede.END)
  emit("moveMemo",memoData)
}

const deleteMome = () =>{
  let memoData:SendMemo = makeSendMemo(ActionCede.DELETE)
  emit("moveMemo",memoData)
}

const makeSendMemo =(codeNumber:number) :SendMemo =>{
  let memo:SendMemo = {
      id: props.data.memo.id,
      text: text.value,
      x:y.value,
      y:x.value,
      actionId:codeNumber,
      boardId:props.boardId
  }
  return memo
}

let x = ref(0)
let y = ref(0)

const colorList:Ref<string[]> = ref([
  "bg-red-400","bg-blue-400","bg-emerald-400"
])

const cloer = () :string =>  {
  return colorList.value[Math.floor( Math.random() * colorList.value.length)]
}

onMounted(() => {
  let kx:number
  let ky:number
  memo.value.style.top = props.data.memo.x + 'px'
  memo.value.style.left = props.data.memo.y + 'px'
  var el:HTMLElement = memo.value;
  watchEffect( () => {
    memo.value.style.top = props.data.memo.x + 'px'
    memo.value.style.left = props.data.memo.y + 'px'
    text.value = props.data.memo.text
    if(props.do){
      opponent.value = props.do
    }else{
      opponent.value = props.do
      memo.value.blur()
    }
  })
  //マウスが要素内で押されたとき、又はタッチされたとき発火
  memo.value.addEventListener("mousedown", mdown)
  // memo.value.addEventListener("touchstart", mdown, false);
  memo.value.addEventListener("dblclick",action);

  //マウスが押された際の関数
  function mdown(e) {

    //タッチデイベントとマウスのイベントの差異を吸収
    if (e.type === "mousedown") {
      var event = e;
    } else {
      var event = e.changedTouches[0];
    }

    //要素内の相対座標を取得
    kx = event.pageX - this.offsetLeft;
    ky = event.pageY - this.offsetTop;
    //ムーブイベントにコールバック
    document.body.addEventListener("mousemove",  mmove);
    document.body.addEventListener("touchmove", mmove);
  }

  //マウスカーソルが動いたときに発火
  async function mmove(e) {
    //ドラッグしている要素を取得
    var drag:Element = await document.getElementsByClassName("drag")[0];
    //同様にマウスとタッチの差異を吸収
    if (e.type === "mousemove") {
      var event = e;
    } else {
      var event = e.changedTouches[0];
    }
    start()
    //フリックしたときに画面を動かさないようにデフォルト動作を抑制
    e.preventDefault();
    //マウスが動いた場所に要素を動かす
    y.value = event.pageY - ky
    x.value = event.pageX - kx
    memo.value.style.top = y.value + "px";
    memo.value.style.left =x.value + "px";

    move() //値を送信

    //マウスボタンが離されたとき、またはカーソルが外れたとき発火
    memo.value.addEventListener("mouseup", mup, false);
    document.body.addEventListener("mouseleave", mup, false);
    memo.value.addEventListener("touchend", mup, false);
    document.body.addEventListener("touchleave", mup, false);
  }

  //マウスボタンが上がったら発火
  async function mup(e) {
      //ムーブベントハンドラの消去
      document.body.removeEventListener("mousemove", mmove, false);
      // drag.removeEventListener("mouseup", mup, false);
      document.body.removeEventListener("touchmove", mmove, false);
      // drag.removeEventListener("touchend", mup, false);
      //クラス名 .drag も消す
      memo.value.classList.remove("drag");
      end()
  }

  function action(e) {
    start()
    view.value = !view.value;
  }
});

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
