<script setup lang="ts">
import Timeline from "@/components/Timeline.vue"
import { get } from "@/utils/http/axios";
import { Ref } from "vue";

type ArticleType = {
  "ID": Number,
  "CreatedAt": String,
  "UpdatedAt": String,
  "DeletedAt": null,
  "title": String,
  "cid": Number,
  "desc": String,
  "content": String,
  "img": String
}

type ArticleListType = Array<ArticleType>
let timelineDesc: Ref<ArticleListType> = ref([])
  
onBeforeMount(async () => {
  let res = await get({ url: '/article/timeline' })
  console.log(res)
  timelineDesc.value = res.list

})
</script>

<template>
  <Timeline :timelineDesc="timelineDesc" class="w-full" />
</template>


<style type="scss">

</style>