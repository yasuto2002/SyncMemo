import type {BoardHistory}  from '../../repository/respons/boardList';
import type { Ref } from 'vue'
import { string } from 'yup';
export default defineNuxtPlugin(() => {
    return {
        provide: {
            getBoards: async(mail:string) : Promise<Ref<BoardHistory[]>> => {
                const boards:Ref<Array<BoardHistory>> = ref([])
                const router = useRouter();
                const config = useRuntimeConfig()
                const { data, pending, refresh, error }  = await useFetch(`${config.apiServer}/board/list`, { method: 'POST',headers: {
                        'Content-Type': 'application/json',
                        'authorization':mail,
                },});
                if(error.value){
                    console.log(error)
                    router.push("/error")
                }
                boards.value = data.value as Array<BoardHistory>
                refresh();
                return boards
            }
        }
    }
})
