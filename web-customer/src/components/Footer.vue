<template>
  <div style="margin-top: 2rem" class="footer">
    <el-row :gutter="10" style="background-color: #54880A">
      <el-col :span="8" class="column-1">
        <h3 class="footer-title">Các phòng chức năng</h3>
        <a href="https://tchc.dlu.edu.vn/">
          <p>Phòng Tổ chức hành chính</p>
        </a>
        <a href="https://pqldt.dlu.edu.vn/">
          <p>Phòng Quản lý Đào tạo</p>
        </a>
        <a href="https://pctsv.dlu.edu.vn/">
          <p>Phòng Chính trị và Công tác Sinh viên</p>
        </a>
        <a href="https://pktkd.dlu.edu.vn/">
          <p>Phòng Quản lý Chất lượng</p>
        </a>
        <a href="https://pkhht.dlu.edu.vn/">
          <p>Phòng Quản lý Khoa học – Hợp tác Quốc tế</p>
        </a>
        <a href="https://ptt.dlu.edu.vn/">
          <p>Phòng Thanh Tra</p>
        </a>
        <a href="https://ptc.dlu.edu.vn/">
          <p>Phòng Tài Chính</p>
        </a>
        <a href="https://pcsvc.dlu.edu.vn/">
          <p>Phòng Cơ sở Vật chất</p>
        </a>
        <a href="https://sdh.dlu.edu.vn/">
          <p>Phòng Quản lý Đào tạo Sau Đại học</p>
        </a>
        <a href="https://ptctt.dlu.edu.vn/">
          <p>Phòng Tạp chí và Truyền Thông</p>
        </a>
      </el-col>
      <el-col :span="8" class="column-2">
        <h3 class="footer-title">Các khoa</h3>
        <a href="http://ktt.dlu.edu.vn/">
          <p>Khoa Toán – Tin</p>
        </a>
        <a href="http://cntt.dlu.edu.vn/">
          <p>Khoa Công nghệ Thông tin</p>
        </a>
        <a href="http://vl.dlu.edu.vn/">
          <p>Khoa Vật lý và Kỹ thuật hạt nhân</p>
        </a>
        <a href="http://khhmt.dlu.edu.vn/">
          <p>Khoa Hóa học và Môi trường</p>
        </a>
        <a href="http://ksh.dlu.edu.vn/">
          <p>Khoa Sinh học</p>
        </a>
        <a href="http://knl.dlu.edu.vn/">
          <p>Khoa Nông lâm</p>
        </a>
        <a href="http://knl.dlu.edu.vn/">
          <p>Khoa Ngữ văn và Lịch sử</p>
        </a>
        <a href="http://nvvh.dlu.edu.vn/">
          <p>Khoa Kinh tế – Quản trị Kinh doanh</p>
        </a>
        <a href="http://kktqt.dlu.edu.vn/">
          <p>Khoa Du lịch</p>
        </a>
        <a href="http://klh.dlu.edu.vn/">
          <p>Khoa Luật học</p>
        </a>
        <a href="http://nn.dlu.edu.vn/">
          <p>Khoa Ngoại ngữ</p>
        </a>
        <a href="http://kqth.dlu.edu.vn/">
          <p>Khoa Quốc tế học</p>
        </a>
        <a href="http://kctxh.dlu.edu.vn/">
          <p>Khoa Xã hội học và Công tác xã hội</p>
        </a>
        <a href="http://sp.dlu.edu.vn/">
          <p>Khoa Sư phạm</p>
        </a>
        <a href="http://kllct.dlu.edu.vn/">
          <p>Khoa Lý luận Chính trị</p>
        </a>
        <a href="http://kgdtc.dlu.edu.vn/">
          <p>Khoa Giáo dục Quốc phòng – An ninh và Thể chất</p>
        </a>
      </el-col>
      <el-col :span="4" class="column-3">
        <h3 class="footer-title">TRƯỜNG ĐẠI HỌC ĐÀ LẠT</h3>
        <p>1 Phù Đổng Thiên Vương - P8 - TP.Đà Lạt</p>
        <p>0263 3822 246 | Fax: 0263.3823380</p>
        <p>info@dlu.edu.vn</p>
        <iframe
          src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3903.2878659389257!2d108.4420162147601!3d11.954560391528823!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x317112d959f88991%3A0x9c66baf1767356fa!2zVHLGsOG7nW5nIMSQ4bqhaSBI4buNYyDEkMOgIEzhuqF0!5e0!3m2!1svi!2s!4v1597049530993!5m2!1svi!2s"
          width="400" height="450" frameborder="0" style="border:0;" allowfullscreen="" aria-hidden="false" tabindex="0"
          data-dashlane-frameid="5711"></iframe>
      </el-col>

    </el-row>
  </div>
  <div class="copy-right">
    <p>
      Copyright &copy; {{ new Date().getFullYear() }} Trung tâm Công nghệ
      Thông tin Trường Đại học Đà Lạt
    </p>
  </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Footer",
};
</script>

<style scoped>
.footer {
  text-align: left;
}

.footer-title {
  text-transform: uppercase;
  color: #fff;
  font-size: 22px;
}

.footer p {
  color: #fff;
  font-size: 16px;
}

.column-1 {
  margin-left: 5%;
}

.column-3 {
  margin-right: 5%;
}

.copy-right {
  background-color: #fff;
  padding: 8px;
}

.footer a {
  text-decoration: none;
  line-height: 1.4em;
}

.footer a:hover {
  color: #FC7616;
}

@media screen and (max-width: 600px) {
  .el-row {
    display: flex;
    flex-direction: column;
  }

  .el-col {
    max-width: 100%;
  }

  .column-3,
  .column-1 {
    margin: 0%;
  }

  iframe{
    width: 350px;
  }
}
</style>
