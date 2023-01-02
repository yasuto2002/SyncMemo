<template>
  <div class="w-full m-auto">
    <TemplateBoradHead @onChildMethodClick = "onChildMethodClick " />
    <TemplateBoard ref="childRef" />
  </div>
</template>
<script setup lang="ts">
import type { Ref } from 'vue';
const childRef:Ref = ref();

const onChildMethodClick = () => {
  childRef.value.coll();
};
</script>
