<template>
  <div>
    <input type="checkbox" id="my-modal" class="modal-toggle" />
    <div class="modal">
      <div class="modal-box">
        <div class="flex justify-between items-center">
          <h3 class="font-bold text-lg">共有URL</h3>
          <div class="modal-action my-auto">
            <label
              for="my-modal"
              class="text-[14px] text-white bg-[#FFC9EF] p-2 rounded-sm"
              >閉じる!</label
            >
          </div>
        </div>

        <div class="flex justify-around items-center my-4">
          <p
            class="py-1 px-1 bottom-1 border-2 border-black border-solid w-[85%]"
          >
            {{ config.domain + route.fullPath }}
          </p>
          <button @click="copy">copy</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
const route = useRoute()
const config = useRuntimeConfig()
const copy = () => {
  navigator.clipboard.writeText(`${config.domain + route.fullPath}`)
}
</script>
