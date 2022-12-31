import type {createRoom} from '../../repository/respons/createRoom';
export default defineNuxtPlugin(() => {
    return {
        provide: {
                createRoom: async (id:string) :Promise<createRoom | null> => {
                const config = useRuntimeConfig()
                const router = useRouter();
                const { data, pending, refresh, error }  = await useFetch(config.apiServer + `/chatroom/create/${id}`, { method: 'POST' });
                if(typeof error.value === "boolean"){
                    return null
                }
                refresh()
                return data.value as createRoom
            }
        }
    }
})