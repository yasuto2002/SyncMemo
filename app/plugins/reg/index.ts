import type {Registration} from '../../repository/request/registration'
import type { Ref } from 'vue'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
import { errCode } from '~~/repository/errCode'
export default defineNuxtPlugin(() => {
    return {
        provide: {
            reg: async (mail:string,token:string) :Promise<errCode> => {
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
                return statusCode.value
            }
        }
    }
})
const onResponseError = async ({ response }: { response: FetchResponse<Error>},statusCode:Ref<number>) => {
	statusCode.value = response.status
};