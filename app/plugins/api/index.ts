export default defineNuxtPlugin(() => {
  return {
    provide: {
      reqApi: async () => {
        let data = await $fetch("https://jsonplaceholder.typicode.com/users")
        return data
      },
    },
  }
})
