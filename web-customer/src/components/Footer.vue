<template>
  <div style="margin-top: 2rem">
    <el-row :gutter="10" style="background-color: #e9e9e9">
      <el-col :span="8">
        <h3>Các phòng chức năng</h3>
        <p>Phòng Tổ chức hành chính</p>
        <p>Phòng Quản lý Đào tạo</p>
        <p>Phòng Chính trị và Công tác Sinh viên</p>
        <p>Phòng Quản lý Chất lượng</p>
        <p>Phòng Quản lý Khoa học – Hợp tác Quốc tế</p>
        <p>Phòng Thanh Tra</p>
        <p>Phòng Tài Chính</p>
        <p>Phòng Cơ sở Vật chất</p>
        <p>Phòng Quản lý Đào tạo Sau Đại học</p>
        <p>Phòng Tạp chí và Truyền Thông</p>
      </el-col>
      <el-col :span="8">
        <h3>Các khoa</h3>
        <p>Khoa Toán – Tin</p>
        <p>Khoa Công nghệ Thông tin</p>
        <p>Khoa Vật lý và Kỹ thuật hạt nhân</p>
        <p>Khoa Hóa học và Môi trường</p>
        <p>Khoa Sinh học</p>
        <p>Khoa Nông lâm</p>
        <p>Khoa Ngữ văn và Lịch sử</p>
        <p>Khoa Kinh tế – Quản trị Kinh doanh</p>
        <p>Khoa Du lịch</p>
        <p>Khoa Luật học</p>
        <p>Khoa Ngoại ngữ</p>
        <p>Khoa Quốc tế học</p>
        <p>Khoa Xã hội học và Công tác xã hội</p>
        <p>Khoa Sư phạm</p>
        <p>Khoa Lý luận Chính trị</p>
        <p>Khoa Giáo dục Quốc phòng – An ninh và Thể chất</p>
      </el-col>
      <el-col :span="8">
        <h3>Liên hệ</h3>
        <p>1 Phù Đổng Thiên Vương - P8 - TP.Đà Lạt</p>
        <p>0263 3822 246 | Fax: 0263.3823380</p>
        <p>info@dlu.edu.vn</p>
      </el-col>
      <el-col>
        <div
          style="
            background-color: #e9e9e9;
            padding: 10px;
            border-top: 1px solid #e1e1e1;
          "
        >
          <p>
            Copyright &copy; {{ new Date().getFullYear() }} Trung tâm Công nghệ
            Thông tin Trường Đại học Đà Lạt
          </p>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Footer",
};
</script>

<style scoped></style>
