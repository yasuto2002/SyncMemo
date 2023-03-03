<template>
  <div>
    <p>test</p>
    <label for="my-modal" class="btn">open modal</label>

    <!-- Put this part before </body> tag -->
    <input type="checkbox" id="my-modal" class="modal-toggle" />
    <div class="modal">
      <div class="modal-box">
        <h3 class="font-bold text-lg">サーバーURL</h3>
        <div class="flex justify-around items-center my-2">
          <p
            class="
              py-2
              px-10
              bottom-1
              border-2 border-black border-solid
              w-[80%]
            "
          >
            https://yasuto0101.com
          </p>
          <button @click="copy">copy</button>
        </div>

        <div class="modal-action">
          <label for="my-modal" class="btn text-[14px]">閉じる!</label>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
const copy = () => {
  navigator.clipboard.writeText("https://yasuto0101");
};
</script>
