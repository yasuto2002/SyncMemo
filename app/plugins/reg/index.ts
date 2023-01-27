import type {Registration} from '../../repository/request/registration'
import type { Ref } from 'vue'
export default defineNuxtPlugin(() => {
    return {
        provide: {
            reg: async (mail:string,token:string) :Promise<boolean> => {
                const statusCode = ref(0)
                const router = useRouter();
                const config = useRuntimeConfig()
                const http = useHttp()
                const registration:Registration = {token:String(token),mail:mail}
                const { data, pending, refresh, error }  = await useAsyncData(String(`register`), () =>
                    $fetch(
                        `${config.apiServer}/register`,
                        { method: 'POST', body:registration,onResponseError: async (ctx) => {
                                statusCode.value = ctx.response.status;
                                await onResponseError(ctx,statusCode);
                        } }
                    ),
                    {
                        initialCache: false
                    }
                )
                if(statusCode.value != 200){
                    switch (statusCode.value) {
                        case http.value.InternalServerError:
                            return null
                            break
                        case http.value.BadRequest:
                            return false
                            break
                    }
                }
                return true
            }
        }
    }
})
const onResponseError = async (data:any,statusCode:Ref<number>) => {
	statusCode.value = data.response.status
};