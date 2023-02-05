<template>
    <div class="flex items-center justify-around my-[5vw]">
        <p>{{boardInfo.name}}</p><p>{{ formatDate(boardInfo.createdAt) }}</p><Nuxt-link :to='{ path: "board",query: { id: boardInfo.id }}' class="text-[#0000FF]">開く</Nuxt-link><button class="text-[#FF0000]" @click="boardDelete">削除</button>
    </div>
</template>
<script setup lang="ts">
    import type { Ref } from 'vue'
    import type {Boards,BoardHistory} from "../../repository/respons/boardList"
    const emit = defineEmits<{
        (event: "bdDelete", id: string): void;
    }>()
    const props = defineProps({
        boardInfo: { type: Object, required: true },
    });
    const formatDate = (dString:string)=>  {
        const date = new Date(dString)
        return `${(date.getFullYear() + "-" + date.getMonth() + 1) + "-" + date.getDate()}`
    }
    const boardDelete = ()=>{
        emit("bdDelete",props.boardInfo.id)
    }
</script>