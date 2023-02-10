<template>
    <form class="w-[50%] m-auto text-[22px] pb-[9vh] bg-[#F9F9F9] h-[100vh]">
        <h1 class="text-center text-[32px] text-[#B9B9B9]">Login</h1>
        <div class="mt-12">
        <p class="text-[#ff0000] text-[12px]" v-if="errNumber === http.BadRequest">メールアドレスかパスワードが間違っています</p>
        <p class="text-[#ff0000] text-[12px]">{{ errors.mail}}</p>
        <p class="text-[#BCB8B8]">メールアドレス</p>
        <input
            type="mail"
            name=""
            class="
            w-full
            border-[1px] border-[#D9D7D7] border-solid
            outline-none
            p-1
            bg-white
            text-[#6B6666]
            inputFoucus
            "
            style="border-radius: 5px"
            v-model="mail"
        />
        </div>
        <div class="mt-12">
        <p class="text-[#ff0000] text-[12px]">{{ errors.password}}</p>
        <p class="text-[#BCB8B8]">パスワード</p>
        <input
            type="password"
            name=""
            class="
            w-full
            border-[1px] border-[#D9D7D7] border-solid
            outline-none
            p-1
            bg-white
            text-[#6B6666]
            inputFoucus
            "
            style="border-radius: 5px"
            autocomplete="on"
            v-model="password"
        />
        </div>
        <div class="flex justify-center items-center mt-12">
        <button
            class="m-auto bg-white text-[#BCB8B8] py-1 px-6 hoverShadow"
            style="box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px"
            @click="onSubmit"
        >
            login
        </button>
        </div>
    </form>
</template>
<script lang="ts" setup>
import { useField, useForm } from "vee-validate"
import * as yup from "yup"
import type {Login} from "../../repository/respons/login"
import type { Ref } from 'vue'
import type { errCode } from '../../repository/errCode'
const router = useRouter();
const config = useRuntimeConfig()
const validateMes = useValidateMes()
const authStore = useAuthStore()
const {authLogin} = authStore
const { $login } = useNuxtApp()
const http = useHttp()
const schema = yup.object({
    mail: yup.string().max(255,validateMes.value.max).required(validateMes.value.required).email(validateMes.value.mail).matches(validateMes.value.regex, validateMes.value.regexMes).trim(),
    password:yup.string().min(10, validateMes.value.min).max(20, validateMes.value.max).matches(validateMes.value.regex, validateMes.value.regexMes).required(validateMes.value.required),
});
useForm({
    validationSchema: schema,
});
const { errors, meta, handleSubmit } = useForm({
    validationSchema: schema,
    initialValues: {
    mail: "",
    password:"",
    },
});
let login:Login = null
let errNumber:Ref<number> = ref()


const { value: mail } = useField("mail")
const { value: password } = useField("password");

const onSubmit = handleSubmit(async(values) => {
    [login,errNumber.value] = await $login(values.mail,values.password)
    console.log(errNumber.value)
    switch (errNumber.value){
        case http.value.InternalServerError:
            router.push("/error")
            break
        case http.value.BadRequest:
            break
        case http.value.Success:
            console.log("成功");
            authLogin() 
            let token = useCookie<{ token: string}>("token",{maxAge: 3600})
            token.value = {token:login.token}
            let cookieMail = useCookie<{ address: string}>("mail",{maxAge: 3600})
            cookieMail.value = {address:values.mail}
            router.push("/")
    }
});
</script>

<style scoped>
.hoverShadow:hover {
    box-shadow: rgba(39, 33, 33, 0.04) 0px 3px 5px !important;
}
.inputFoucus:focus {
    box-shadow: rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;
}
</style>
