<script lang="ts">
  import day from 'dayjs';

  const columns = [
    {
      title: 'ID',
      dataIndex: 'ID',
      width: '5%',
      key: 'id',
      align: 'center',
    },
    {
      title: '更新日期',
      dataIndex: 'UpdatedAt',
      width: '10%',
      key: 'UpdatedAt',
      align: 'center',
      customRender: (val) => {
        return val ? day(val).format('YYYY年MM月DD日 HH:mm') : '暂无';
      },
    },
    {
      title: '分类',
      dataIndex: 'Category.name',
      width: '5%',
      key: 'name',
      align: 'center',
    },
    {
      title: '文章标题',
      dataIndex: 'title',
      width: '15%',
      key: 'title',
      align: 'center',
    },
    {
      title: '文章描述',
      dataIndex: 'desc',
      width: '20%',
      key: 'desc',
      align: 'center',
    },
    {
      title: '缩略图',
      dataIndex: 'img',
      width: '20%',
      key: 'img',
      align: 'center',
      scopedSlots: { customRender: 'img' },
    },
    {
      title: '操作',
      width: '15%',
      key: 'action',
      align: 'center',
      scopedSlots: { customRender: 'action' },
    },
  ];

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
        Artlist: [],
        Catelist: [],
        columns,
        queryParam: {
          title: '',
          pagesize: 5,
          pagenum: 1,
        },
      };
    },
    created() {
      this.getArtList();
      this.getCateList();
    },
    methods: {
      // 获取文章列表
      async getArtList() {
        const { data: res } = await this.$http.get('/article', {
          params: {
            title: this.queryParam.title,
            pagesize: this.queryParam.pagesize,
            pagenum: this.queryParam.pagenum,
          },
        });
        console.log(res);
        if (res.code !== 200) {
          if (res.code === 1004 || 1005 || 1006 || 1007) {
            window.sessionStorage.clear();
            this.$router.push('/login');
          }
          this.$message.error(res.message);
        }

        this.Artlist = res.list;
        this.pagination.total = res.pagination.total;
      },
      // 获取分类
      async getCateList() {
        const { data: res } = await this.$http.get('category');
        if (res.code !== 200) return this.$message.error(res.message);
        this.Catelist = res.list;
        this.pagination.total = res.total;
      },
      // 更改分页
      handleTableChange(pagination, filters, sorter) {
        var pager = { ...this.pagination };
        pager.current = pagination.current;
        pager.pageSize = pagination.pageSize;
        this.queryParam.pagesize = pagination.pageSize;
        this.queryParam.pagenum = pagination.current;

        if (pagination.pageSize !== this.pagination.pageSize) {
          this.queryParam.pagenum = 1;
          pager.current = 1;
        }
        this.pagination = pager;
        this.getArtList();
      },
      // 删除文章
      deleteArt(id) {
        this.$confirm({
          title: '提示：请再次确认',
          content: '确定要删除该文章吗？一旦删除，无法恢复',
          onOk: async () => {
            const { data: res } = await this.$http.delete(`article/${id}`);
            if (res.status != 200) return this.$message.error(res.message);
            this.$message.success('删除成功');
            this.getArtList();
          },
          onCancel: () => {
            this.$message.info('已取消删除');
          },
        });
      },
      // 查询分类下的文章
      CateChange(value) {
        this.getCateArt(value);
      },
      async getCateArt(id) {
        const { data: res } = await this.$http.get(`article/list/${id}`, {
          params: { pagesize: this.queryParam.pagesize, pagenum: this.queryParam.pagenum },
        });
        if (res.code !== 200) return this.$message.error(res.message);
        this.Artlist = res.list;
        this.pagination.total = res.pagination.total;
      },
    },
  };
</script>

<template>
  <el-card>
    <el-row :gutter="20">
      <el-col :span="6">
        <el-input
          v-model="queryParam.title"
          placeholder="输入文章名查找"
          enter-button
          allowClear
          @search="getArtList"
        />
      </el-col>
      <el-col :span="4">
        <el-button type="primary" @click="$router.push('/addart')">新增</el-button>
      </el-col>

      <el-col :span="3">
        <el-select placeholder="请选择分类" style="width: 200px" @change="CateChange">
          <el-option v-for="item in Catelist" :key="item.id" :value="item.id"
            >{{ item.name }}
          </el-option>
        </el-select>
      </el-col>
      <el-col :span="1">
        <el-button type="info" @click="getArtList()">显示全部</el-button>
      </el-col>
    </el-row>

    <el-table rowKey="ID" :data="Artlist" @change="handleTableChange">
      <el-table-column prop="ID" label="编号" width="100" />
      <el-table-column label="封面" width="180">
        <template #default="scope">
          <img class="h-10 w-10" :src="scope.row.img" />
        </template>
      </el-table-column>
      <el-table-column prop="title" label="标题" width="180"></el-table-column>
      <el-table-column label="操作" width="230">
        <template #default="scope" class="flex justify-evenly">
          <el-button
            size="small"
            type="primary"
            icon="edit"
            style="margin-right: 15px"
            @click="$router.push(`/addart/${data.ID}`)"
            >编辑</el-button
          >
          <el-button
            size="small"
            type="danger"
            icon="delete"
            style="margin-right: 15px"
            @click="deleteArt(data.ID)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<style scoped>
  .actionSlot {
    display: flex;
    justify-content: center;
  }

  .ArtImg {
    height: 100%;
    width: 100%;
  }

  .ArtImg img {
    width: 100px;
    height: 80px;
  }
</style>
