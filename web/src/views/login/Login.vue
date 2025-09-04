<template>
  <div class="login_container">
    <div class="login_box">
      <div class="avatar_box">
        <img src="../assets/logo.png" alt />
      </div>
      <el-form :model="loginForm" :rules="loginRules" ref="loginFormRef" label-width="0px" class="login_form">
        <el-form-item prop="username">
          <el-input v-model="loginForm.username" prefix-icon="iconfont icon-user"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input 
            v-model="loginForm.password" 
            prefix-icon="iconfont icon-3702mima" 
            :type="pwdType"
          >
            <i slot="suffix" class="iconfont icon-showpassword" @click="showPwd"></i>
          </el-input>
        </el-form-item>
        <el-form-item class="btns">
          <el-button type="primary" @click="login">登录</el-button>
          <el-button type="info" @click="resetLoginForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loginForm: {
        username: 'linglong', // 默认测试用户名
        password: '123456'    // 默认测试密码
      },
      pwdType: 'password',
      loginRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 3, max: 16, message: '长度在 3 到 16 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    showPwd() {
      this.pwdType = this.pwdType === 'password' ? 'text' : 'password'
    },
    resetLoginForm() {
      this.$refs.loginFormRef.resetFields()
    },
    async login() {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return

        try {
          // 后端地址（需与 Docker 中后端服务端口匹配，示例为 8080）
          const baseURL = 'http://你的服务器IP:8080' 
          this.$http.defaults.baseURL = baseURL
          
          // 发送登录请求（form-data 格式与后端匹配）
          const { data: resp } = await this.$http.post('/api/v1/auth', this.loginForm, {
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
          })

          if (resp.code === 200) {
            // 存储 Token（核心修复：前端保存 Token 用于后续认证）
            window.sessionStorage.setItem('token', resp.data.token)
            this.$message.success('登录成功')
            this.$router.push('/home') // 跳转到首页
          } else {
            this.$message.error(resp.msg || '登录失败')
          }
        } catch (error) {
          this.$message.error('服务器连接失败，请检查服务')
          console.error('登录错误:', error)
        }
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login_container {
  background: url(../assets/img/bg.png);
  height: 100%;
}
.login_box {
  width: 450px;
  height: 300px;
  background-color: #fff;
  border-radius: 10px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
.avatar_box {
  width: 130px;
  height: 130px;
  border-radius: 50%;
  border: 1px solid #eee;
  padding: 10px;
  box-shadow: 0 0 10px #ddd;
  position: absolute;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;
  img {
    width: 100%;
    height: 100%;
    border-radius: 50%;
    background-color: #eee;
  }
}
.login_form {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
  .icon-showpassword {
    margin-right: 8px;
    cursor: pointer;
  }
}
.btns {
  display: flex;
  justify-content: flex-end;
}
</style>
