import type {makeBoardRes} from '../../repository/respons/makeBoard';
import type { Ref } from 'vue'
import type { errCode } from '../../repository/errCode'
import type { Make } from '../../repository/request/make'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
export default defineNuxtPlugin(() => {
return {
provide: {
    gestMakeBoard: async (name:string,pass:string) :Promise<[makeBoardRes,errCode]> => {
    const router = useRouter();
    const config = useRuntimeConfig()
    const statusCode:Ref<errCode> = ref(200);
    const make:Make = {name:name,password:pass}
    const { data, pending, refresh, error }  = await useAsyncData<makeBoardRes>(String(`gestMakeBoard`), () =>
                $fetch(
                    `${config.apiServer}/gestMakeBoard`,
                    { method: 'POST'
                        ,body:make,onResponseError: async (ctx) => {
                            statusCode.value = ctx.response.status;
                            await onResponseError(ctx,statusCode);
                    } }
                ),
                {
                    initialCache: false
                }
            )
    if(typeof error.value === "boolean"){
        return null
    }
    return [data.value,statusCode.value]
    }
}
}
})
const onResponseError = async ({ response }: { response: FetchResponse<Error>},statusCode:Ref<number>) => {
	statusCode.value = response.status
}