<template>
    <div class="m-auto pb-12">
        <div>
            <h1 class="text-center text-[20px]">作成済みボード</h1>
        </div>
        <div class="w-[60%] m-auto">
            <TemplateBoardInfo v-if="length != 0"  v-for="(board ,i) in viewBoard" :key="i" :boardInfo="board"/>
        </div>
        <TemplatePagination >
            <PartsPaginationBtn class="bg-white" v-for="j in (length % defoCount)" :class="{'bg-[#F9FAF7] hover:cursor-default':  pageCount == j }"   @click="pageCalc(j)" >{{ String(j) }}</PartsPaginationBtn>
        </TemplatePagination>
    </div>
</template>
<script setup lang="ts">
    import type {Boards,BoardHistory} from "../../repository/respons/boardList"
    import type { Ref } from 'vue'
    const router = useRouter()
    const { $getBoardsAll } = useNuxtApp()
    const token = useCookie<{ token: string}>("token")
    const length:Ref<number> = ref(0)
    const boards:Ref<BoardHistory[]> = ref(new Array)
    const viewBoard:Ref<BoardHistory[]> = ref(new Array)
    const boardListSet = async()=>{
        boards.value = await $getBoardsAll(token.value.token)
        length.value = boards.value.length
        arryCalc()
    }
    boardListSet()

    const defoCount:Ref<number> = ref(6) //ボードの数
    const pageCount:Ref<number> = ref(1) //ページの位置

    const arryCalc = () =>{
        if(boards.value.length < 6){ //ボードの数が6個以下
            viewBoard.value = boards.value.slice(0, boards.value.length-1)
        }else{
            viewBoard.value = boards.value.slice(0, 6)
        }
    }
    
    //ページネーション
    const pageCalc = (num:number) =>{
        if(num === pageCount.value){ //同じページ
            return
        }
        else if(num == 1 && num < 0){
            arryCalc()
        }
        else if(((num * defoCount.value) -1) < boards.value.length){ //ボードの数が６の倍数に満たない時
            viewBoard.value = boards.value.slice((num-1) * defoCount.value, boards.value.length-1)
        }else{
            viewBoard.value = boards.value.slice((num-1) * defoCount.value, (num * defoCount.value) -1)
        }
        pageCount.value = num
    }
    
</script>
