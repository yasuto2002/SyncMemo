<template>
  <div class="w-full m-auto">
    <div class="flex justify-between w-[40%] mb-[1%] ml-[5%]">
      <PartsBoardName>3限算数の時間</PartsBoardName>
      <PartsMakeName>yasuto</PartsMakeName>
      <button @click="onChildMethodClick" class="text-[#D8CFCF]">
        メモ追加
      </button>
    </div>
    <templateBoard ref="childRef" />
  </div>
</template>
<script setup>
const childRef = ref();

const onChildMethodClick = () => {
  childRef.value.coll();
};
</script>
