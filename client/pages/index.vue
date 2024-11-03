<script lang="ts" setup>
const value = ref('fdkljdkdfdfjkf')
</script>

<template>
  <MonacoEditor v-model="value" lang="C" :style="{width: '100%', height: '100%'}" />
</template>