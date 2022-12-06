<template>
    <div class="w-[47%] bg-white h-full p-[5%]">
        <h1 class="text-center  text-[20px] mb-[3vw]">作成済みボード</h1>
        <div class="flex items-center justify-around my-10" v-for="i of 4" :key="i">
            <p>算数の時間①</p>
            <p>2020-7-15</p>
            <p class="text-[#0000FF] hover:cursor-pointer">開く</p>
            <p class="text-[#FF0000] hover:cursor-pointer">削除</p>
        </div>
        <div class="flex justify-end w-full"><NuxtLink to="/" class="text-[#CACCD0] border-b-[1px] border-[#CACCD0] border-solid">もっと見る</NuxtLink></div>
    </div>
</template>