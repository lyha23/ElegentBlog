<template>
  <el-card>
    <el-row :gutter="20">
      <el-col :span="4">
        <el-button type="primary" @click="addCateVisible = true">新增分类</el-button>
      </el-col>
    </el-row>

    <el-table rowKey="id" :data="Catelist" @change="handleTableChange">
      <el-table-column label="名称" prop="name" />

      <el-table-column label="操作">
        <template #default="scope">
          <el-button type=" primary" icon="edit" style="margin-right: 15px" @click="editCate(data)">编辑</el-button>
          <el-button type="danger" icon="delete" style="margin-right: 15px" @click="deleteCate(data)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>

  <!-- 新增分类区域 -->
  <el-dialog closable title="新增分类" v-model="addCateVisible" width="60%" @ok="addCateOk" @cancel="addCateCancel"
    destroyOnClose>
    <el-form :model="newCate" :rules="addCateRules" ref="addCateRef">
      <el-form-item label="分类名称" prop="name">
        <el-input v-model="newCate.name" />
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="addCateVisible = false">Cancel</el-button>
        <el-button type="primary" @click="addCateOk">
          OK
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 编辑分类区域 -->
  <el-dialog closable destroyOnClose title="编辑分类" v-model="editCateVisible" width="60%" @ok="editCateOk"
    @cancel="editCateCancel">
    <el-form :model="CateInfo" :rules="CateRules" ref="addCateRef">
      <el-form-item label="分类名称" prop="name">
        <el-input v-model="CateInfo.name" />
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="editCateVisible = false">Cancel</el-button>
        <el-button type="primary" @click="editCateOk">
          OK
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts">
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id',
    align: 'center',
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center',
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      Catelist: [],
      CateInfo: {
        name: '',
        id: 0,
      },
      newCate: {
        name: '',
      },
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1,
      },
      editVisible: false,
      CateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.CateInfo.name === '') {
                callback(new Error('请输入分类名'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      addCateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newCate.name === '') {
                callback(new Error('请输入分类名'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      editCateVisible: false,
      addCateVisible: false,
    }
  },
  created() {
    this.getCateList()
  },
  methods: {
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('category', {
        pagesize: this.queryParam.pagesize,
        pagenum: this.queryParam.pagenum,
      })

      if (res.code !== 200) {
        if (res.code === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.Catelist = res.list
    },
    // 更改分页
    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getCateList()
    },
    // 删除分类
    deleteCate(data) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该分类吗？一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`category/${data.id}`)
          if (res.code != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getCateList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    // 新增分类
    addCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('category/add', {
          name: this.newCate.name,
        })
        if (res.code != 200) return this.$message.error(res.message)
        this.$refs.addCateRef.resetFields()
        this.addCateVisible = false
        this.$message.success('添加分类成功')
        await this.getCateList()
      })
    },
    addCateCancel() {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
      this.$message.info('新增分类已取消')
    },
    // 编辑分类
    async editCate(data) {
      this.editCateVisible = true
      this.CateInfo = data
      this.CateInfo.id = data.id
    },
    editCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`category/${this.CateInfo.id}`, {
          name: this.CateInfo.name,
        })
        if (res.code != 200) return this.$message.error(res.message)
        this.editCateVisible = false
        this.$message.success('更新分类信息成功')
        this.getCateList()
      })
    },
    editCateCancel() {
      this.$refs.addCateRef.resetFields()
      this.editCateVisible = false
      this.$message.info('编辑已取消')
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>