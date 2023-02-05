import type {BoardHistory}  from '../../repository/respons/boardList'
import type { Ref } from 'vue'
import { string } from 'yup'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
import { errCode } from '~~/repository/errCode'

export default defineNuxtPlugin(() => {
    return {
        provide: {
            getBoardsAll: async(token:string) : Promise<[BoardHistory[],errCode]> => {
                const statusCode = ref(0)
                const boards:Ref<Array<BoardHistory>> = ref([])
                const router = useRouter();
                const config = useRuntimeConfig()
                const { data, pending, refresh, error }  = await useAsyncData<BoardHistory[]>(String(`getBoardsAll`), () =>
                    $fetch(
                        `${config.apiServer}/board/listAll`,
                        { method: 'POST',onResponseError: async (ctx) => {
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
                return [data.value,statusCode.value]
            }
        }
    }
})
const onResponseError = async ({ response }: { response: FetchResponse<Error>},statusCode:Ref<number>) => {
	statusCode.value = response.status
}