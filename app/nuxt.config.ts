import { defineNuxtConfig } from "nuxt"

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  // ssr: false,
  // target: "static",
  modules: [
    // Using package name (recommended usage)
  ],
  css: ["@/assets/css/reset.css", "@/assets/css/main.css"],

  ssr: false,
  // 追加
  build: {
    postcss: {
      postcssOptions: {
        plugins: {
          tailwindcss: {},
          autoprefixer: {},
        },
      },
    },
  },
  runtimeConfig: {
    // The private keys which are only available within server-side
    apiSecret: process.env.APISERVER,
    // Keys within public, will be also exposed to the client-side
    public: {
      apiServer: process.env.APISERVER,
      apiContainer: process.env.APICONTAINER,
      socket: process.env.SOCKET,
      domain: process.env.DOMAIN,
    },
  },
  alias: {
    yup: "yup/lib/index.js",
  },
})
