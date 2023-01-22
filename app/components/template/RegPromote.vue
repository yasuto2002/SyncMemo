<template>
    <div class="w-[47%] bg-white h-full p-[5%] relative">
        <h1 class="text-center  text-[20px] mb-[3vw] opacity-20">作成済みボード</h1>
        <div class="flex items-center justify-around my-10 opacity-20" v-for="i in 4">
            <p>ボードの名前</p>
            <p>2020-7-15</p>
            <NuxtLink to="/" class="text-[#0000FF] hover:cursor-pointer" tabindex="-1" style="pointer-events: none">開く</NuxtLink>
            <p class="text-[#FF0000] hover:cursor-pointer">削除</p>
        </div>
        <div class="flex justify-end w-full"><NuxtLink to="/" class="text-[#CACCD0] border-b-[1px] border-[#CACCD0] border-solid opacity-20" style="pointer-events: none">もっと見る</NuxtLink></div>
        <div class="w-[100%] h-[100%] absolute top-0 left-0 flex justify-center items-center">
            <h1 class="text-[1.5vw]">会員登録するとボードを保存できます</h1>
        </div>
    </div>
</template>
<script setup lang="ts">
</script>