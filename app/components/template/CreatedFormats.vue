<template>
    <form class="bg-white p-[5%] w-[47%] rounded-[10px] h-full">
        <h1 class="text-center text-[20px]"><slot name="formatName" /></h1>
        <div class="w-[80%] m-auto">
            <p class="mt-[3vw]">ボードの名前</p>
            <p class="text-[#ff0000] text-[12px]">{{ errors.name}}</p>
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
            v-model="name"
            /><br />
            <input type="checkbox" id="validity" v-model="checked" :disabled="!loginStatus"/>
            <label for="validity" :class="{'opacity-60':!loginStatus}">パスワードを有効にする</label>
            <p class="text-[#ff0000]" v-if="!loginStatus">ログインユーザーのみパスワードを有効にできます</p>
            <p class="mt-[3vw]" v-bind:class= "{'text-[#F3F5F4]' : !checked || !loginStatus}">パスワード</p>
            <p class="text-[#ff0000] text-[12px]">{{ errors.password}}</p>
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
            v-bind:class= "{'border-[#ACA5A5]' : checked && loginStatus}"
            :disabled="!checked || !loginStatus"
            v-model="password"
            ref="input"
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
            @click="onSubmit"
            >
            作る
            </button>
        </div>
    </form>
</template>
<script setup lang="ts">
import { useField, useForm } from "vee-validate"
import * as yup from "yup"
import makeBoard from '~~/plugins/makeBoard'
import type {makeBoardRes} from '../../repository/respons/makeBoard'
import type { Ref } from 'vue'
const router = useRouter()
const authStore = useAuthStore()
const checked = useState('ref1-key', () => false)
const { $makeBoard } = useNuxtApp()
const { $gestMakeBoard } = useNuxtApp()
const { authState } = authStore
const loginStatus = ref(computed(() => authState.value))
const input:Ref<HTMLElement> = ref(null)
const validateMes = useValidateMes()
const schema = yup.object({
    name: yup.string().max(255,validateMes.value.max).required(validateMes.value.required).matches(validateMes.value.regex, validateMes.value.regexMes).trim(),
    password:yup.string().max(20, validateMes.value.max).matches(validateMes.value.regex, validateMes.value.regexMes),
});
useForm({
    validationSchema: schema,
})
const { errors, meta, handleSubmit } = useForm({
        validationSchema: schema,
        initialValues: {
        name:"",
        password:"",
    },
})
const { value: name } = useField("name")
const { value: password } = useField("password")
const onSubmit = handleSubmit(async(values) => {
    const { $makeBoard } = useNuxtApp()
    let id:makeBoardRes = null
    if(authState.value){
        const token = useCookie<{ token: string}>("token")
        id = await $makeBoard(values.name,values.password,token.value.token)
    }else{
        id = await $gestMakeBoard(values.name,"")
    }
    if(id === null){
        router.push("/error")
        return
    }
    router.push({ path: 'board',query: { id: id.id }})
})
onMounted(() => {
    watchEffect(() => {
        if(!loginStatus.value){
            password.value = ""
        }
    })
})
</script>
