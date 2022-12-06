<template>
    <form class="bg-white p-[5%] w-[47%] rounded-[10px] h-full">
        <h1 class="text-center text-[20px]"><slot name="formatName" /></h1>
        <div class="w-[80%] m-auto">
            <p class="mt-[3vw]">ボードの名前</p>
            <input
            type="text"
            name=""
            id=""
            class="
            bg-white
            outline-none
            border-[1px] border-[#ACA5A5] border-solid
            w-full
            mb-[3vw]
            rounded-[10px]
            h-10
            "
            /><br />
            <input type="checkbox" id="validity" v-model="checked"/>
            <label for="validity">パスワードを有効にする</label>
            <p class="mt-[3vw]" v-bind:class= "{'text-[#F3F5F4]' : !checked}">パスワード</p>
            <input
            type="password"
            name=""
            id=""
            class="
                bg-white
                outline-none
                border-[1px] border-solid
                w-full
                mb-[3vw]
                rounded-[10px]
                h-10
            "
            v-bind:class= "{'border-[#ACA5A5]' : checked}"
            :disabled="!checked"
        />
        </div>
        <div class="w-full flex justify-center">
            <button
            class="
                w-[30%]
                m-auto
                bg-[#BCB8B8]
                p-1
                text-center
                rounded-[10px]
                mx-auto
                mt-3
                text-white
            "
            type="button"
            >
            作る
            </button>
        </div>
    </form>
</template>
<script setup lang="ts">
const checked = useState('ref1-key', () => false)
</script>
