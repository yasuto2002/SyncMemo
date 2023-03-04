type CounterState = {
  count: number
}
const useAuthState = () => {
  return useState("authState", () => false)
}

const useAuthLogin = () => {
  const authState = useAuthState()
  authState.value = true
}

const useAuthLogout = () => {
  const authState = useAuthState()
  authState.value = false
}

// export const useGetAuthState = () =>{
//     return useState('getAuthState', () => {
//         const authState = useAuthState()
//         return authState
//     })
// }

export const useAuthStore = () => {
  const authState = useAuthState()
  return {
    authState: authState,
    authLogin: useAuthLogin,
    authLogout: useAuthLogout,
  }
}
