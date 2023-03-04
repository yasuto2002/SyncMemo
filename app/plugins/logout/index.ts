import type { Ref } from "vue"
import type { errCode } from "../../repository/errCode"
import type { FetchContext, FetchResponse } from "ohmyfetch"
export default defineNuxtPlugin(() => {
  return {
    provide: {
      tokenDelete: async (token: string): Promise<errCode> => {
        const router = useRouter()
        const config = useRuntimeConfig()
        const statusCode: Ref<errCode> = ref(200)
        const { data, pending, refresh, error } = await useAsyncData(
          String(`logout`),
          () =>
            $fetch(`${config.apiServer}/logout`, {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
                Authorization: token,
              },
              onResponseError: async (ctx) => {
                statusCode.value = ctx.response.status
                await onResponseError(ctx, statusCode)
              },
            }),
          {
            initialCache: false,
          }
        )
        if (typeof error.value === "boolean") {
          return statusCode.value
        }
        return statusCode.value
      },
    },
  }
})
const onResponseError = async (
  { response }: { response: FetchResponse<Error> },
  statusCode: Ref<number>
) => {
  statusCode.value = response.status
}
