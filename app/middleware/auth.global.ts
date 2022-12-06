export default defineNuxtRouteMiddleware(async(to, from) => {
    const config = useRuntimeConfig()
    const logToken = useCookie('logToken')
    const authStore = useAuthStore()
    const {authLogin} = authStore
    if(logToken){
        const {data}= await useFetch(`${config.apiContainer}/loginCheck`,
        { method: 'POST', body: { 
            token:logToken.value
        }}
        )
        if(!data.value){
        }else{
            let flgData:any = data
            if (flgData.value.flg) {
                authLogin()
                console.log("login")
            }else{
                console.log("loginしてない")
            }
        }
    }
})