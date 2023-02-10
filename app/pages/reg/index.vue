<template>
  <form class="w-[50%] m-auto text-[22px] pb-[9vh]">
    <h1 class="text-center text-[32px] text-[#B9B9B9]">会員登録</h1>
    <div class="mt-12">
      <p class="text-[#ff0000] text-[12px]">{{ errors.name}}</p>
      <p class="text-[#BCB8B8]">名前</p>
      <input
        type="text"
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
        v-model="name"
      />
    </div>
    <div class="mt-12">
      <p class="text-[#ff0000] text-[12px]" v-if="code===http.BadRequest">メールアドレスがすでに使われています</p>
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
    <div class="mt-12">
      <p class="text-[#ff0000] text-[12px]">{{ errors.conPassword}}</p>
      <p class="text-[#BCB8B8]">パスワード確認</p>
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
        v-model="conPassword"
      />
    </div>
    <div class="flex justify-center items-center mt-12">
      <button
        class="m-auto bg-white text-[#BCB8B8] py-1 px-6 hoverShadow"
        style="box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px"
        @click="onSubmit"
      >
        登録
      </button>
    </div>
  </form>
</template>
<script lang="ts" setup>
import { useField, useForm } from "vee-validate"
import * as yup from "yup"
import type { Ref } from 'vue'
import { errCode } from "../../repository/errCode"
const router = useRouter()
const config = useRuntimeConfig()
const validateMes = useValidateMes()
const { $casual } = useNuxtApp()
const route = useRoute()
const authStore = useAuthStore()
const {authState} = authStore
const {authLogin} = authStore

const schema = yup.object({
  mail: yup.string().max(255,validateMes.value.max).required(validateMes.value.required).email(validateMes.value.mail).matches(validateMes.value.regex, validateMes.value.regexMes).trim(),
  name: yup.string().max(10,validateMes.value.max).min(2,validateMes.value.min).required(validateMes.value.required).matches(validateMes.value.regex, validateMes.value.regexMes),
  password:yup.string().min(10, validateMes.value.min).max(20, validateMes.value.max).matches(validateMes.value.regex, validateMes.value.regexMes).required(validateMes.value.required),
  conPassword:yup.string().min(10, validateMes.value.min).max(20, validateMes.value.max).matches(validateMes.value.regex, validateMes.value.regexMes).required(validateMes.value.required).oneOf([yup.ref("password")],validateMes.value.match),
});
useForm({
  validationSchema: schema,
});
const { errors, meta, handleSubmit } = useForm({
  validationSchema: schema,
  initialValues: {
    mail: "",
    name:"",
    password:"",
    conPassword:""
  },
});
const { value: mail } = useField("mail")
const { value: name } = useField("name")
const { value: password } = useField("password");
const { value: conPassword } = useField("conPassword");
const errMes = ref()
const http = useHttp()
const result:Ref<boolean> = ref(null)
let code:Ref<errCode> = ref()
const onSubmit = handleSubmit(async(values) => {
  code.value = await $casual(values.name,values.mail,values.password)
  console.log(code.value)
  switch (code.value){
      case http.value.InternalServerError:
          router.push("/error")
          break
      case http.value.BadRequest:
          break
      case http.value.Unauthorized:
          router.push("/login")
          break
      default:
      const mail = useCookie<{ address: string}>("mail",{maxAge: 3600})
      mail.value = {address:values.mail}
      router.push({ name: 'reg-token', params: { mail: values.mail } })
  }
  refresh()
})
const refresh = () => refreshNuxtData('code')
</script>

<style scoped>
.hoverShadow:hover {
  box-shadow: rgba(39, 33, 33, 0.04) 0px 3px 5px !important;
}
.inputFoucus:focus {
  box-shadow: rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;
}
</style>
