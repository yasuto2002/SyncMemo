<template>
  <div class="w-full">
    <div
      class="flex items-center justify-between w-[90%] m-auto pb-11 h-[40vw]"
    >
      <TemplateCreatedFormats>
        <template v-slot:formatName>ボードを作る</template>
      </TemplateCreatedFormats>
      <TemplateBoardHistory v-if="authState" />
      <TemplateRegPromote v-if="!authState"></TemplateRegPromote>
    </div>
  </div>
</template>
<script setup>
const authStore = useAuthStore()
const { authState } = authStore
</script>
