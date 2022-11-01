<template>
  <div class="w-full">
    <el-card>
      <h3>{{ id ? '编辑文章' : '新增文章' }}</h3>
      <el-form :model="artInfo" ref="artInfoRef" :rules="artInfoRules" :hideRequiredMark="true">
        <el-row :gutter="24">
          <el-col :span="16">
            <el-form-item label="文章标题" prop="title">
              <el-input style="width: 300px" v-model="artInfo.title" />
            </el-form-item>
            <el-form-item label="文章描述" prop="desc">
              <el-input type="textarea" v-model="artInfo.desc" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="文章分类" prop="cid">
              <el-select style="width: 200px" v-model="artInfo.cid" placeholder="请选择分类" @change="cateChange">
                <el-option v-for="item in Catelist" :key="item.id" :value="item.id">
                  {{
                      item.name
                  }}
                </el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="文章缩略图" prop="img">
              <el-upload listType="picture" name="file" :action="upUrl" :headers="headers" @change="upChange">
                <el-button>
                  <el-icon type="upload" />点击上传
                </el-button>

                <template v-if="id">
                  <img :src="artInfo.img" style="width: 120px; height: 100px; margin-left: 15px" />
                </template>
              </el-upload>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="文章内容" prop="content">
          <v-md-editor v-model="artInfo.content" :disabled-menus="[]" height="600px" @upload-image="handleUploadImage">

          </v-md-editor>
        </el-form-item>

        <el-form-item>
          <el-button type="danger" style="margin-right: 15px" @click="artOk(artInfo.id)">
            {{
                artInfo.id ? '更新' : '提交'
            }}
          </el-button>
          <el-button type="primary" @click="addCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script lang="ts" >
import VMdEditor from "@kangc/v-md-editor";
export default {
  components: { VMdEditor },
  props: ['id'],
  data() {
    return {
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: '',
      },
      Catelist: [],
      upUrl: this.$http.baseURL + 'upload',
      headers: {},
      fileList: [],
      artInfoRules: {
        title: [{ required: true, message: '请输入文章标题', trigger: 'change' }],
        cid: [{ required: true, message: '请选择文章分类', trigger: 'change' }],
        desc: [
          { required: true, message: '请输入文章描述', trigger: 'change' },
          { max: 120, message: '描述最多可写120个字符', trigger: 'change' },
        ],
        // img: [{ required: true, message: '请选择文章缩略图', trigger: 'change' }],
        content: [{ required: true, message: '请输入文章内容', trigger: 'change' }],
      },
    }
  },
  mounted() {
    this.getCateList()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    // 查询文章信息
    async getArtInfo(id) {
      const { data: res } = await this.$http.get(`admin/article/info/${id}`)
      if (res.code !== 200) {
        if (res.code === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$message.error(res.message)
          this.$router.push('/login')
          return
        }
        this.$message.error(res.message)
        return
      }
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('category')
      if (res.code !== 200) return this.$message.error(res.message)
      this.Catelist = res.list
    },
    // 选择分类
    cateChange(value) {
      this.artInfo.cid = value
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
    // 上传图片
    // upChange(info) {
    //   if (info.file.status !== 'uploading') {
    //   }
    //   if (info.file.status === 'done') {
    //     this.$message.success(`图片上传成功`)
    //     this.artInfo.img = info.file.response.url
    //   } else if (info.file.status === 'error') {
    //     this.$message.error(`图片上传失败`)
    //   }
    // },
    // 提交&&更新文章
    // artOk(id) {
    //   this.$refs.artInfoRef.validate(async (valid) => {
    //     if (!valid) return this.$message.error('参数验证未通过，请按要求录入文章内容')
    //     if (id === 0) {
    //       const { data: res } = await this.$http.post('article/add', this.artInfo)
    //       if (res.code !== 200) return this.$message.error(res.message)
    //       this.$router.push('/artlist')
    //       this.$message.success('添加文章成功')
    //     } else {
    //       const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
    //       if (res.code !== 200) return this.$message.error(res.message)
    //       this.$router.push('/artlist')
    //       this.$message.success('更新文章成功')
    //     }
    //   })
    // },

    // addCancel() {
    //   this.$refs.artInfoRef.resetFields()
    // },
  },
}
</script>
