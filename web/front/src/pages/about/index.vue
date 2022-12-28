<template>
  <v-md-preview :text="userStore.aboutMe"></v-md-preview>
  <!-- <VisTopology :id="'VisTopology'" /> -->
</template>
<script lang="ts" setup>
  import { useUserStore } from '@/store';
  const userStore = useUserStore();
</script>
