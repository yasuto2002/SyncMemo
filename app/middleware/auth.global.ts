import type { Ref } from "vue"
import type { errCode } from "../repository/errCode"
export default defineNuxtRouteMiddleware(async (to, from) => {
  const config = useRuntimeConfig()
  const logToken = useCookie("logToken")
  const authStore = useAuthStore()
  const { authLogin } = authStore
  const token = useCookie<{ token: string }>("token")
  const router = useRouter()
  const statusCode: Ref<errCode> = ref(200)
  const http = useHttp()
  if (typeof token.value != "undefined") {
    const http = useHttp()
    const { data, pending, refresh, error } = await useAsyncData(
      String(`loginCheck`),
      () =>
        $fetch(`${config.apiServer}/loginCheck`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            authorization: token.value.token,
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
    if (
      error.value ||
      statusCode.value == http.value.Unauthorized ||
      statusCode.value == http.value.InternalServerError
    ) {
      return
    } else if (statusCode.value == http.value.Success) {
      authLogin()
    }
  }
})
const onResponseError = async (data: any, statusCode: Ref<number>) => {
  statusCode.value = data.response.status
}
