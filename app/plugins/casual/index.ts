import type {Casual} from '../../repository/request/casual'
import type { Ref } from 'vue'
import { errCode } from '../../repository/errCode'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
export default defineNuxtPlugin(() => {
    return {
        provide: {
            casual: async (name:string,mail:string,pass:string) :Promise<errCode> => {
                const statusCode:Ref<errCode> = ref(200);
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
                return statusCode.value
            }
        }
    }
})
const onResponseError = async ({ response }: { response: FetchResponse<Error>},statusCode:Ref<number>) => {
	statusCode.value = response.status
}