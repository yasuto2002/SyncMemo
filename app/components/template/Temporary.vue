<template>
  <div
    class="w-[1200px] h-[1000px] bg-[#F3F3F3] m-auto"
    style="overflow: hidden; position: relative; box-sizing: border-box"
  >
    <div
      class="drag-and-drop"
      id="red-box"
      style="
        width: 200px;
        height: 200px;
        background-color: blue;
        top: 200px;
        white-space: pre-wrap;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 2;
        overflow: hidden;
      "
      ref="root"
      v-for="memo in memos"
      :key="memo"
      v-bind:style="{ top: memo.y + 'px', left: memo.x + 'px' }"
    >
      {{ memo.text }}
    </div>
  </div>
</template>
<script setup lang="ts">
const root = ref(null)
const config = useRuntimeConfig()
let memos = ref()
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
</style>
