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
            v-model="boardName"
            /><br />
            <input type="checkbox" id="validity" v-model="checked" :disabled="!loginStatus"/>
            <label for="validity" :class="{'opacity-60':!loginStatus}">パスワードを有効にする</label>
            <p class="text-[#ff0000]" v-if="!loginStatus">ログインユーザーのみパスワードを有効にできます</p>
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
            v-model="boardPassword"
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
            @click="make"
            >
            作る
            </button>
        </div>
    </form>
</template>
<script setup lang="ts">
import makeBoard from '~~/plugins/makeBoard';
import type {makeBoardRes} from '../../repository/respons/makeBoard';
import type { Ref } from 'vue';
const boardName:Ref<string> = ref("")
const boardPassword:Ref<string> = ref("")
const router = useRouter()
const authStore = useAuthStore()
const checked = useState('ref1-key', () => false)
const { $makeBoard } = useNuxtApp()
const { authState } = authStore
const loginStatus = ref(computed(() => authState.value))
const make = async() =>{
    const { $makeBoard } = useNuxtApp()
    const id:makeBoardRes = await $makeBoard(boardName.value,boardPassword.value)
    if(id === null){
        router.push("/error")
        return
    }
    router.push({ path: 'board',query: { id: id.id }})
}

</script>
