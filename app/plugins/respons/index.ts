import type {BoardHistory}  from '../../repository/respons/boardList'
import type { Ref } from 'vue'
import { string } from 'yup'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
import { errCode } from '~~/repository/errCode'

export default defineNuxtPlugin(() => {
    return {
        provide: {
            getBoards: async(token:string) : Promise<[BoardHistory[],errCode]> => {
                const statusCode:Ref = ref(200)
                const boards:Ref<Array<BoardHistory>> = ref([])
                const router = useRouter();
                const config = useRuntimeConfig()
                const { data, pending, refresh, error }  = await useAsyncData<BoardHistory[]>(String(`getBoards`), () =>
                    $fetch(
                        `${config.apiServer}/board/list`,
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
