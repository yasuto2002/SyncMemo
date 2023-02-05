<template>
    <div class="flex justify-between w-[40%] mb-[1%] ml-[5%]">
        <PartsBoardName>3限算数の時間</PartsBoardName>
        <button @click="onClick" class="text-[#D8CFCF]">
        メモ追加
        </button>
    </div>
</template>
<script setup lang="ts">
const emits = defineEmits<{(e: 'onChildMethodClick'): void}>()
const onClick = (): void => {
    emits('onChildMethodClick')
}
</script>