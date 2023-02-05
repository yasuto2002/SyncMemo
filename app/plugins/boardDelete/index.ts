import type { boardDelete }  from '../../repository/request/boardDelete'
import type { Ref } from 'vue'
import { string } from 'yup'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
import type { errCode } from '../../repository/errCode'

export default defineNuxtPlugin(() => {
    return {
        provide: {
            boardDelete: async(token:string,id:string) : Promise<errCode> => {
                const statusCode:Ref<errCode> = ref(0)
                const router = useRouter();
                const config = useRuntimeConfig()
                const boardDelete:boardDelete = { boardId : id}
                const { data, pending, refresh, error }  = await useAsyncData(String(`boardDelete`), () =>
                    $fetch(
                        `${config.apiServer}/board/delete`,
                        { method: 'POST',body:boardDelete,onResponseError: async (ctx) => {
                                statusCode.value = ctx.response.status;
                                await onResponseError(ctx,statusCode);
                        },headers: {
                            'Content-Type': 'application/json',
                            'authorization':token,
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
                await refresh()
                return statusCode.value
            }
        }
    }
})
const onResponseError = async ({ response }: { response: FetchResponse<Error>},statusCode:Ref<number>) => {
	statusCode.value = response.status
}