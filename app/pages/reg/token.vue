<template>
    <div class="pb-[5vw]">
        <div class="w-[40%] m-auto text-center text-[25px] text-[#BCB8B8] border-[1px] border-[#D9D7D7] border-solid p-[3vw] shadow-inner" style="border-radius: 5px">
            <h1 class="mb-[2vw]">認証コードを送信しました。</h1>
            <p class="mb-[2vw]">メールアドレス <span>{{mail}}</span> の認証を行います。</p>
            <p class="mb-[2vw]">送信された4桁の認証用コードを入力してください</p>
            <form action="
            ">
                <h1 class="mb-[2vw]">認証コード(4桁)</h1>
                <p class="text-[#ff0000] text-[12px]">{{errors.token}}</p>
                <p class="text-[#ff0000] text-[12px]" v-if="result === false">コードが間違っています</p>
                <input
                    type="number"
                    name=""
                    class="
                    w-[90%]
                    border-[1px] border-[#D9D7D7] border-solid
                    outline-none
                    p-1
                    bg-white
                    text-[#6B6666]
                    inputFoucus
                    mx-auto
                    mb-[2vw]
                    "
                    style="border-radius: 5px"
                    autocomplete="on"
                    max="9999"
                    min="1111"
                    v-model="token"
                />
                <button class="bg-white py-[3%] px-[5%] text-[20px] hover:shadow-none shadow-sm" @click="onSubmit">認証を行う</button>
            </form>
        </div>
    </div>
</template>
<script lang="ts" setup>
    import { Result } from "postcss";
    import { useField, useForm } from "vee-validate"
    import * as yup from "yup"
    import type { Ref } from 'vue'
    import { errCode } from "../../repository/errCode"
    const validateMes = useValidateMes()
    const config = useRuntimeConfig()
    const authStore = useAuthStore()
    const {authLogin} = authStore
    const {authState} = authStore
    const { $reg } = useNuxtApp()
    const router = useRouter()
    const http = useHttp()
    const schema = yup.object({
    token: yup.number().max(9999,validateMes.value.max).required(validateMes.value.required).min(1111,"桁少ない").nullable()
        .transform((value, originalValue) =>
        String(originalValue).trim() === '' ? null : value
        )
    });
    useForm({
        validationSchema: schema,
    });
    const { errors, meta, handleSubmit } = useForm({
        validationSchema: schema,
        initialValues: {
            token: "",
        },
    })
    const route = useRoute()
    const mail:string = route.params.mail as string
    if(typeof mail == "undefined"){
        router.push("/")
    }

    const { value: token } = useField("token")
    const result:Ref<boolean> = ref(true)
    // const regData = await useCookie('regData',60 * 60)
    const onSubmit = handleSubmit(async(values) => {
        let code:errCode
        code = await $reg(mail,values.token)
        switch (code){
            case http.value.InternalServerError:
                router.push("/error")
                break
            case http.value.BadRequest:
                router.push("/reg")
                break
            case http.value.Unauthorized:
                router.push("/login")
                break
            default:
                router.push("/login")
                alert("完了！")
        }
    })
</script>
