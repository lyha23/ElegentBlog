<script setup>
  import { wxlogin } from '/@/api/user/index';
  import { useRouter } from 'vue-router';
  import { setToken } from '/@/utils/auth';
  import { useUserStore } from '/@/store/index';
  import router from '/@/router';

  let { currentRoute } = useRouter();

  const userStore = useUserStore();

  // console.log(currentRoute.value.query.code)
  wxlogin({ code: currentRoute.value.query.code })
    .then((res) => {
      console.log(res, '成功');
      setToken(res.data.Token);
      userStore.setInfo(res.data.userInfo);
      router.push('/');
    })
    .catch((e) => {
      console.log(e, '失败');
    });
</script>

<template> </template>
