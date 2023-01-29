import type { Memos } from '../../repository/request/memos'
import type { Memo as ResMemos} from '../../repository/respons/memo';
import type { Ref } from 'vue'
import type { errCode } from '../../repository/errCode'
import type {FetchContext,FetchResponse} from 'ohmyfetch'
export default defineNuxtPlugin(() => {
    return {
        provide: {
            memos: async (id:string) :Promise<[Array<ResMemos>,errCode]> => {
                const statusCode:Ref<errCode> = ref(200);
                const router = useRouter();
                const config = useRuntimeConfig()
                const http = useHttp()
                const reqMemos:Memos = {id:id}
                const { data, pending, refresh, error }  = await useAsyncData<Array<ResMemos>>(String(`memos`), () =>
                    $fetch(
                        `${config.apiServer}/memos`,
                        { method: 'POST', body:reqMemos,onResponseError: async (ctx) => {
                                statusCode.value = ctx.response.status;
                                await onResponseError(ctx,statusCode);
                        } }
                    ),
                    {
                        initialCache: false
                    }
                )
                if(statusCode.value != http.value.Success){
                    switch (statusCode.value) {
                        case http.value.InternalServerError:
                            return [null,statusCode.value]
                        
                        case http.value.BadRequest:
                            return [null,statusCode.value]
                    }
                }else{
                    return [data.value,statusCode.value]
                }
            }
        }
    }
})
const onResponseError = async ({ response }: { response: FetchResponse<Error>},statusCode:Ref<number>) => {
	statusCode.value = response.status
}