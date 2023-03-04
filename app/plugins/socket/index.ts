import { io, Socket } from "socket.io-client"
interface position {
  id: string
  text: string
  x: number
  y: number
}
interface returnData {
  socket: Socket
  make: Function
  send: Function
}
export default defineNuxtPlugin(() => {
  return {
    provide: {
      makeSoket: (): returnData => {
        const config = useRuntimeConfig()
        const socket = io(config.apiServer)
        const sendMemo = (memoData: position) => {
          socket.emit("moveMemo", memoData)
        }
        const makeMemo = (): void => {
          socket.emit("makeMemo")
        }
        return { socket: socket, make: makeMemo, send: sendMemo }
      },
    },
  }
})
