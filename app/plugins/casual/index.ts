import type {Casual} from '../../repository/request/casual'
import type { Ref } from 'vue'
export default defineNuxtPlugin(() => {
    return {
        provide: {
            casual: async (name:string,mail:string,pass:string) :Promise<boolean> => {
                const statusCode = ref(200);
                const router = useRouter();
                const config = useRuntimeConfig()
                const http = useHttp()
                const registration:Casual = {name:name,password:pass,mail:mail}
                const { data, pending, refresh, error }  = await useAsyncData(String(`casual`), () =>
                    $fetch(
                        `${config.apiServer}/casual`,
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
                            break;
                        
                        case http.value.BadRequest:
                            return false
                    }
                }else{
                    return true
                }
            }
        }
    }
})
const onResponseError = async (data:any,statusCode:Ref<number>) => {
	statusCode.value = data.response.status
}