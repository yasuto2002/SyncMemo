import type { Login } from '../../repository/request/login'
import type { Login as ResLogin} from '../../repository/respons/login';
import type { Ref } from 'vue'
import type { errCode } from '../../repository/errCode'
export default defineNuxtPlugin(() => {
    return {
        provide: {
            login: async (mail:string,pass:string) :Promise<[ResLogin,errCode]> => {
                const statusCode:Ref<errCode> = ref(200);
                const router = useRouter();
                const config = useRuntimeConfig()
                const http = useHttp()
                const log:Login = {mail:mail,password:pass}
                const { data, pending, refresh, error }  = await useAsyncData(String(`login`), () =>
                    $fetch(
                        `${config.apiServer}/login`,
                        { method: 'POST', body:log,onResponseError: async (ctx) => {
                                statusCode.value = ctx.response.status;
                                await onResponseError(ctx,statusCode);
                        } }
                    ),
                    {
                        initialCache: false
                    }
                )
                return [data.value as ResLogin,statusCode.value]
            }
        }
    }
})
const onResponseError = async (data:any,statusCode:Ref<number>) => {
	statusCode.value = data.response.status
}