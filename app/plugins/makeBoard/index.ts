import { string } from "yup"

export default defineNuxtPlugin(() => {
    return {
      provide: {
        makeBoard: async () => {
          let {ID}:any = await $fetch("http://localhost:8080/makeBoard");
          return ID;
        }
      }
    }
  })