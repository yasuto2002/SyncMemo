<template>
    <div class="m-auto pb-12">
        <div>
            <h1 class="text-center text-[20px]">作成済みボード</h1>
        </div>
        <div class="w-[60%] m-auto">
            <TemplateBoardInfo  v-for="i in 6" :key="i"/>
        </div>
        <div class="w-full">
            <div class="btn-group m-auto flex justify-center py-9">
                    <button class="btn bg-white border-white">1</button>
                    <button class="btn  bg-white border-white">2</button>
                    <button class="btn bg-white border-white">3</button>
                    <button class="btn bg-white border-white">4</button>
            </div>
        </div>
    </div>
    
    
</template>
<style scoped>
    .border-white{
        border-color: #fff !important;
    }
</style>
