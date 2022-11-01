<template>
  <el-card>
    <el-form labelAlign="left">

      <el-form-item label="作者名称">
        <el-input style="width: 300px" v-model="profileInfo.name"></el-input>
      </el-form-item>

      <el-form-item label="个人简介">
        <el-input type="textarea" v-model="profileInfo.desc" />
      </el-form-item>

      <el-form-item label="QQ号码">
        <el-input style="width: 300px" v-model="profileInfo.qq_chat" />
      </el-form-item>

      <el-form-item label="微信">
        <el-input style="width: 300px" v-model="profileInfo.wechat" />
      </el-form-item>

      <el-form-item label="微博">
        <el-input style="width: 300px" v-model="profileInfo.weibo" />
      </el-form-item>

      <el-form-item label="B站地址">
        <el-input style="width: 300px" v-model="profileInfo.bili" />
      </el-form-item>

      <el-form-item label="Github">
        <el-input style="width: 300px" v-model="profileInfo.github" />
      </el-form-item>

      <el-form-item label="Email">
        <el-input style="width: 300px" v-model="profileInfo.email" />
      </el-form-item>

      <el-form-item label="头像">
        <el-upload listType="picture" name="file" :action="upUrl" :headers="headers" @change="avatarChange">
          <el-button style="margin-right:10px">
            <el-icon type="upload" />点击上传
          </el-button>

          <template v-if="profileInfo.avatar">
            <img :src="profileInfo.avatar" style="width: 120px; height: 100px" />
          </template>
        </el-upload>
      </el-form-item>

      <el-form-item label="个人简介">
        <v-md-editor v-model="profileInfo.aboutMe" :disabled-menus="[]" height="600px"
          @upload-image="handleUploadImage">
        </v-md-editor>
      </el-form-item>

      <el-form-item>
        <el-button type="danger" style="margin-right: 15px" @click="updateProfile">更新</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script lang="ts">

export default {
  data() {
    return {
      profileInfo: {
        id: 1,
        name: '',
        desc: '',
        qq_chat: '',
        wechat: '',
        weibo: '',
        bili: '',
        github: '',
        email: '',
        avatar: '',
        aboutMe: '',
      },
      upUrl: this.$http.baseURL + 'upload',
      headers: {},
    }
  },
  created() {
    this.getProfileInfo()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
  },
  methods: {
    // 获取个人设置
    async getProfileInfo() {
      const { data: res } = await this.$http.get(`profile/${this.profileInfo.id}`)
      if (res.code !== 200) {
        if (res.code === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      console.log(res.data)
      this.profileInfo = res.data
    },

    // 上传头像
    avatarChange(info) {
      if (info.file.status !== 'uploading') {
      }
      if (info.file.status === 'done') {
        this.$message.success(`图片上传成功`)
        const imgUrl = info.file.response.url
        this.profileInfo.avatar = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`图片上传失败`)
      }
    },

    async handleUploadImage(event, insertImage, files) {
      // 拿到 files 之后上传到文件服务器，然后向编辑框中插入对应的内容
      console.log(files);
      console.log(event)
      let form = new FormData();
      let file = event.target.files[0]
      form.append('file', file);
      let config = {
        headers: { 'Content-Type': 'multipart/form-data' }
      }
      let { data: res } = await this.$http.post('upload', form, config)
      console.log(res)
      // 此处只做示例
      insertImage({
        url: res.url,
        desc: res.name,
        // width: 'auto',
        // height: 'auto',
      });
    },

    // 更新
    async updateProfile() {
      const { data: res } = await this.$http.put(`profile/1`, this.profileInfo)
      if (res.code !== 200) return this.$message.error(res.message)
      this.$message.success(`个人信息更新成功`)
      this.$router.push('/index')
    },
  },
}
</script>

<style scoped>
.upBtn {
  margin-right: 10px;
}
</style>
