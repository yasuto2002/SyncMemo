import type {createRoom} from '../../repository/respons/createRoom';
import { errCode } from '~~/repository/errCode'
import type { Ref } from 'vue'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
export default defineNuxtPlugin(() => {
    return {
        provide: {
            createRoom: async(id:string) : Promise<[createRoom,errCode]> => {
                const statusCode = ref(200)
                const router = useRouter();
                const config = useRuntimeConfig()
                const { data, pending, refresh, error }  = await useAsyncData<createRoom>(String(`createRoom`), () =>
                    $fetch(
                        config.apiServer + `/chatroom/create/${id}`,
                        { method: 'POST',onResponseError: async (ctx) => {
                                statusCode.value = ctx.response.status;
                                await onResponseError(ctx,statusCode);
                        },headers: {
                            'Content-Type': 'application/json',
                    } }
                    ),
                    {
                        initialCache: false
                    }
                )
                if(error.value){
                    console.log(error)
                    router.push("/error")
                }
                refresh()
                return [data.value,statusCode.value]
            }
        }
    }
})
const onResponseError = async ({ response }: { response: FetchResponse<Error>},statusCode:Ref<number>) => {
	statusCode.value = response.status
}