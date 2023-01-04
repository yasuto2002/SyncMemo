<template>
    <div class="m-auto pb-12">
        <div>
            <h1 class="text-center text-[20px]">作成済みボード</h1>
        </div>
        <div class="w-[60%] m-auto">
            <TemplateBoardInfo  v-for="i in 6" :key="i"/>
        </div>
        <TemplatePagination />
    </div>
    
    
</template>
