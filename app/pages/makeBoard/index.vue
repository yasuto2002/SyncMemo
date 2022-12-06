<template>
  <div class="w-full">
    <h1 class="text-center pb-4 text-[20px]">Hello User!!!!!</h1>
    <div class="flex items-center justify-between w-[90%] m-auto pb-11 h-[40vw]">
      <PartsCreatedFormats>
        <template v-slot:formatName>ボードを作る</template>
      </PartsCreatedFormats>
      <PartsBoardHistory />
    </div>
  </div>
</template>
