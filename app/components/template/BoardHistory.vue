<template>
    <div class="w-[47%] bg-white h-full p-[5%]">
        <h1 class="text-center  text-[20px] mb-[3vw]">作成済みボード</h1>
        <div class="flex items-center justify-around my-10" v-for="(board,i) in boards">
            <p class="w-[25%]">{{board.name}}</p>
            <p class="w-[25%]">{{ formatDate(board.createdAt) }}</p>
            <NuxtLink :to='{ path: "board",query: { id: board.id }}' class="text-[#0000FF] hover:cursor-pointer w-[25%] text-center">開く</NuxtLink>
            <p class="text-[#FF0000] hover:cursor-pointer w-[25%] text-center" @click="bdDelete(board.id)">削除</p>
        </div>
        <div class="flex justify-end w-full"><NuxtLink to="/myBoards" class="text-[#CACCD0] border-b-[1px] border-[#CACCD0] border-solid mr-9">もっと見る</NuxtLink></div>
    </div>
</template>
<script setup lang="ts">
    import type {Boards,BoardHistory} from "../../repository/respons/boardList"
    import { onBeforeRouteLeave, onBeforeRouteUpdate } from 'vue-router'
    import { errCode } from "../../repository/errCode";
    import type { Ref } from 'vue'
    const router = useRouter()
    const { $getBoards } = useNuxtApp()
    const { $boardDelete } = useNuxtApp()
    const token = useCookie<{ token: string}>("token")
    const boards:Ref<BoardHistory[]> = ref(new Array)
    const authStore = useAuthStore()
    const { authState } = authStore
    const http = useHttp()
    const boardListSet = async()=>{
        boards.value = await $getBoards(token.value.token)
    }
    boardListSet()
    const formatDate = (dString:string)=>  {
        const date = new Date(dString)
        return `${(date.getFullYear() + "-" + date.getMonth() + 1) + "-" + date.getDate()}`
    }
    const bdDelete = async(id:string) =>{
        const token = useCookie<{ token: string}>("token")
        if(typeof token.value != "undefined" && authState.value){
            let code:errCode =  await $boardDelete(token.value.token,id)
            switch (code){
                case http.value.InternalServerError:
                    router.push("/error")
                    break
                case http.value.Unauthorized:
                    router.push("/error")
                    break
                case http.value.Unauthorized:
                    router.push("/login")
                    break
                default:
                boardListSet()
            }
        }else{
            router.push("/login")
        }
    } 
</script>