<script setup lang="ts">
import { useArticleStore } from '@/store';
import { get } from '/@/utils/http/axios';
import router from '@/router'
const Route = useRoute()
const articleState = useArticleStore()

let state = reactive({
  category: null as unknown as Number,
  list: [],
  pagination: {
    pageSize: 10,
    page: 1,
    totalPage: 10,
  }
})

state.category = Route.query.category as unknown as number


async function updateList() {
  state.pagination.page = Number(Route.query.page)
  let url = Route.query.category ? `article/list/${Route.query.category}` : 'article'
  let res = await getArtList(url)
  articleState.setInfo(res)
  state.pagination = res.pagination
  state.list = res.list
}

async function getArtList(url: string) {
  let res = await get({
    url: url,
    params: {
      pagesize: state.pagination.pageSize,
      pagenum: state.pagination.page
    }
  })
  return res
}

router.beforeEach(async (to, _) => {
  if (to.name == "page") {
    state.pagination.page = Number(to.query.page)
    let url = to.query.category ? `article/list/${to.query.category}` : 'article'
    let res = await getArtList(url)
    articleState.setInfo(res)
    state.pagination = res.pagination
    state.list = res.list
  }
})

onBeforeMount(async () => {
  await updateList()
})

</script>
  
  
<template>
  <div v-for="(item ) in state.list">
    <cover :item="item"></cover>
  </div>
  <footer class="flex mt-5">
    <nav class="flex nav">
      <router-link :to="`/page?page=${state.pagination.page-1}`" v-show="state.pagination.page>1" class=" navItem ">
        <i class="iconfont icon-zuo"></i>
      </router-link>
      <div v-show="state.pagination.page==1" class=" navItem ">
        <i class="iconfont icon-zuo"></i>
      </div>
      <div class="flex " v-if="state.pagination.totalPage<10">
        <router-link v-for="(item,_) in state.pagination.totalPage" :to="`/page?page=${item}`" class="navItem"
          :class="item==Number(Route.query.page)?'selectedNav':''">
          {{item}}
        </router-link>
      </div>
      <div v-else class="flex">
        <router-link :to="`/page/?page=1`" class="navItem" :class="1==Number(Route.query.page)?'selectedNav':''">
          1
        </router-link>
        <router-link v-show="[1,2,3].indexOf(Number(state.pagination.page))!=-1" v-for="(item,index) in [2,3]"
          :to="`/page/?page=${item}`" class="navItem" :class="item==Number(Route.query.page)?'selectedNav':''">
          {{item}}
        </router-link>
        <span v-show="Number(Route.query.page) >3" class="navItem">
          ...
        </span>
        <router-link
          v-show="[1,2,3,state.pagination.totalPage-2,state.pagination.totalPage-1,state.pagination.totalPage-2,state.pagination.totalPage].indexOf(Number(Route.query.page))==-1"
          v-for="(item,index) in [state.pagination.page-1,state.pagination.page,state.pagination.page+1]"
          :to="`/page/?page=${item}`" class="navItem" :class="item==Number(Route.query.page)?'selectedNav':''">
          {{item}}
        </router-link>
        <span v-show="Number(Route.query.page) <state.pagination.totalPage-2" class="navItem">
          ...
        </span>
        <router-link
          v-show="[state.pagination.totalPage-2,state.pagination.totalPage-1,state.pagination.totalPage].indexOf(Number(state.pagination.page))!=-1"
          v-for="(item,index) in [state.pagination.totalPage-2,state.pagination.totalPage-1]"
          :to="`/page/?page=${item}`" class="navItem" :class="item==Number(Route.query.page)?'selectedNav':''">
          {{item}}
        </router-link>
        <router-link :to="`/page/?page=${state.pagination.totalPage}`" class="navItem"
          :class="state.pagination.totalPage==Number(Route.query.page)?'selectedNav':''">
          {{state.pagination.totalPage}}
        </router-link>
      </div>
      <router-link :to="`/page/?page=${state.pagination.page+1}`"
        v-show="state.pagination.page<state.pagination.totalPage" class=" navItem ">
        <i class="iconfont icon-you"></i>
      </router-link>
      <div :to="`/page/?page=${state.pagination.page+1}`" v-show="state.pagination.page==state.pagination.totalPage"
        class=" navItem ">
        <i class="iconfont icon-you"></i>
      </div>
    </nav>
  </footer>

</template>
  
  
  
<style lang="scss" scoped>
.nav {
  @apply rounded-xl bg-[#f5f6f5];

  .navItem {
    @apply flex font-medium h-9 opacity-90 text-[#999] w-9 items-center justify-center;
  }

  .selectedNav {
    @apply rounded-xl bg-[#6495ed] shadow-sm text-light-50;
  }
}


.btn:active::after {
  transform: translate(-50%, -50%) scale(0);
  opacity: 0.3;
  transition: 0s;
}
</style>
  