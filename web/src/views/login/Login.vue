<template>
  <div class="login_container">
    <div class="login_box">
      <!-- 头部区域 -->
      <div class="avatar_box">
        <img src="../assets/logo.png" alt />
      </div>
      <!-- 表单区域 -->
      <el-form :model="loginForm" :rules="loginRules" ref="loginFormRef" label-width="0px" class="login_form">
        <!-- 用户名 -->
        <el-form-item prop="username">
          <el-input v-model="loginForm.username" prefix-icon="iconfont icon-user"></el-input>
        </el-form-item>

        <!-- 密码 -->
        <el-form-item prop="password">
          <el-input 
            v-model="loginForm.password" 
            prefix-icon="iconfont icon-3702mima" 
            :type="pwdType"
          >
            <i slot="suffix" class="iconfont icon-showpassword" @click="showPwd"></i>
          </el-input>
        </el-form-item>

        <!-- 按钮 -->
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
        username: '',
        password: ''
      },
      pwdType: 'password', // 密码显示类型
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
          // 发送登录请求（表单提交方式，与后端formData接收匹配）
          const { data: resp } = await this.$http.post('/api/v1/auth', this.loginForm, {
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
          })

          if (resp.code !== 200) {
            return this.$message.error(resp.msg || '登录失败')
          }

          // 保存token并跳转
          window.sessionStorage.setItem('token', resp.data.token)
          this.$message.success('登录成功')
          this.$router.push('/home')
        } catch (error) {
          this.$message.error('服务器连接失败，请检查服务是否运行')
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
