<template>
    <div class="w-[47%] bg-white h-full p-[5%]">
        <h1 class="text-center  text-[20px] mb-[3vw]">作成済みボード</h1>
        <div class="flex items-center justify-around my-10" v-for="board in boards" :key="board">
            <p>算数の時間①</p>
            <p>2020-7-15</p>
            <router-link :to='{ path: "board",query: { id: board._id }}' class="text-[#0000FF] hover:cursor-pointer">開く</router-link>
            <p class="text-[#FF0000] hover:cursor-pointer">削除</p>
        </div>
        <div class="flex justify-end w-full"><NuxtLink to="/" class="text-[#CACCD0] border-b-[1px] border-[#CACCD0] border-solid">もっと見る</NuxtLink></div>
    </div>
</template>
<script setup>
    const router = useRouter();
    const boards = ref([])
    const { data, pending, refresh, error }  = await useFetch("http://localhost:8080/boardList", { method: 'POST',headers: {
            'Content-Type': 'application/json',
          },});
    if(error.value){
        console.log(error)
        // router.push("/error")
        // return
    }
    let jData = JSON.parse(data.value)
    boards.value = jData.boards
    refresh();
  watch(data, (newData) => {
    let jData = JSON.parse(data.value)
    boards.value = jData.boards
    console.log(boards)
  })
</script>