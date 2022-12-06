<template>
  <div>
    <p
      class="text-[#D8CFCF]"
      v-for="(item, i) in list"
      :ref="
        (el) => {
          if (el) divs[i] = el;
        }
      "
      :key="i"
      :data="i"
    >
      作成者:{{ item }}<slot />
    </p>
  </div>
</template>
<script setup>
const test = ref([]);
const list = reactive([1]);
const divs = ref([]);
onMounted(() => {
  console.log(divs.value[0].style);
});
onBeforeUpdate(() => {
  divs.value = [];
});
const click = () => {
  list.push(1);
};
</script>
