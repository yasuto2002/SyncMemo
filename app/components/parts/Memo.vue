<template>
  <div
    class="drag-and-drop bg-[#fffdfd] text-[#7c4e4e] rounded-[10px] elv"
    id="red-box"
    ref="memo"
  >
    <p v-if="view" class="w-full h-full p-[5%]">{{ data.text }}</p>
    <textarea
      name=""
      id=""
      v-model="data.text"
      v-if="!view"
      class="resize-none focus:outline-none w-full h-full p-[5%]"
      @blur="out"
      @keydown="input"
    ></textarea>
  </div>
</template>
<script setup>
const memo = ref(null);
// let text = ref("");
let view = ref(true);
// const divs = ref([]);
const out = () => {
  view.value = !view.value;
  let memoData = {
      id: props.data.memo.id,
      text: props.data.memo.text,
      x:x.value,
      y:y.value,
  };
  emit("moveMemo",memoData)
};
const input = () =>{
  let memoData = {
      id: props.data.memo._id,
      text: props.data.memo.text,
      x:y.value,
      y:x.value,
  };
  emit("moveMemo",memoData)
}
const props = defineProps({
  data: { type: Object, required: true },
});
let x = ref(0)
let y = ref(0)
const emit = defineEmits();
onMounted(() => {
  var kx;
  var ky;
  memo.value.style.top = props.data.memo.x + 'px'
  memo.value.style.left = props.data.memo.y + 'px'
  var el = memo;
  watchEffect(() => {
    memo.value.style.top = props.data.memo.x + 'px'
    memo.value.style.left = props.data.memo.y + 'px'
  })
  //マウスが要素内で押されたとき、又はタッチされたとき発火
  memo.value.addEventListener("mousedown", mdown)
  // memo.value.addEventListener("touchstart", mdown, false);
  memo.value.addEventListener("dblclick",action);

  //マウスが押された際の関数
  function mdown(e) {
    //クラス名に .drag を追加
    // obj.classList.add("drag");

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
    var drag = await document.getElementsByClassName("drag")[0];
    //同様にマウスとタッチの差異を吸収
    if (e.type === "mousemove") {
      var event = e;
    } else {
      var event = e.changedTouches[0];
    }

    //フリックしたときに画面を動かさないようにデフォルト動作を抑制
    e.preventDefault();
    //マウスが動いた場所に要素を動かす
    y.value = event.pageY - ky
    x.value = event.pageX - kx
    memo.value.style.top = y.value + "px";
    memo.value.style.left =x.value + "px";
    let memoData = {
      id: props.data.memo._id,
      text: props.data.memo.text,
      x:y.value,
      y:x.value,
    };
    emit("moveMemo",memoData)
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
  }
  function action(e) {
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
