import type {makeBoardRes} from '../../repository/respons/makeBoard';
export default defineNuxtPlugin(() => {
  return {
    provide: {
      makeBoard: async (name:string,pass:string) :Promise<makeBoardRes | null> => {
        const router = useRouter();
        const config = useRuntimeConfig()
        const { data, pending, refresh, error }  = await useFetch(`${config.apiServer}/makeBoard`, { method: 'POST', body: {Name : name,Password:pass} });
        if(typeof error.value === "boolean"){
            return null
        }
        return data.value as makeBoardRes
      }
    }
  }
})