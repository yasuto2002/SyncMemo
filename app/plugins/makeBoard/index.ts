import type {makeBoardRes} from '../../repository/respons/makeBoard';
import type { Ref } from 'vue'
import type { errCode } from '../../repository/errCode'
import type { Make } from '../../repository/request/make'
export default defineNuxtPlugin(() => {
  return {
    provide: {
      makeBoard: async (name:string,pass:string,token:string) :Promise<makeBoardRes | null> => {
        const router = useRouter();
        const config = useRuntimeConfig()
        const statusCode:Ref<errCode> = ref(200);
        const make:Make = {name:name,password:pass}
        const { data, pending, refresh, error }  = await useAsyncData(String(`makeBoard`), () =>
                    $fetch(
                        `${config.apiServer}/board/make`,
                        { method: 'POST', headers: {
                          'Content-Type': 'application/json',
                          'authorization':token,}
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
        return data.value as makeBoardRes
      }
    }
  }
})
const onResponseError = async (data:any,statusCode:Ref<number>) => {
	statusCode.value = data.response.status
}