<template>
    <el-form
    label-position="top"
    label-width="auto"
    :model="form"
    class="form"
    >
    <el-form-item label="账号">
      <el-input v-model="form.account" clearable  placeholder="请输入邮箱/手机号" maxlength="30"/>
    </el-form-item>
    <el-form-item label="密码">
      <el-input type="password" v-model="form.pass" clearable placeholder="请输入密码" maxlength="30" show-password/>
    </el-form-item>
    <el-form-item>
        <el-col :span="15">
            <el-button type="primary" @click="submit()" size="large" color="#626aef">
                 Login
            </el-button>
        </el-col>
        <el-col :span="8">
            <el-link href="/signup">create account?</el-link>
        </el-col>
    </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
    import { reactive, ref } from 'vue'
    import type { FormInstance } from 'element-plus'
    import { mobileSignin } from '@/api/mobile'
    import { emailSignin } from '@/api/email'
    import { useTokenStore } from '@/stores/token'
    const formRef = ref<FormInstance>()
    const form = reactive({
        account: '',
        pass: '',
    })
    function submit() {
        if(form.account.includes('@')){
            emailSignin(form).then((res) => {
                const r = res.data
                if (r.code == 0) {
                    useTokenStore().token = r.data.token
                    ElNotification({
                        showClose: true,
                        message: "登录成功",
                        type: 'success',
                    })
                    window.location.href = "/userhome";
                }else {
                    ElNotification({
                        showClose: true,
                        message: r.message,
                        type: 'error',
                    })
                }
            })
        }
        else {
            mobileSignin(form).then((res) => {
                const r = res.data
                if (r.code == 0) {
                    useTokenStore().token = r.data.token
                    ElNotification({
                        showClose: true,
                        message: "登录成功",
                        type: 'success',
                    })
                    window.location.href = "/userhome";
                }else {
                    ElNotification({
                        showClose: true,
                        message: r.message,
                        type: 'error',
                    })
                }
            })
        }
    }
</script>
<style scoped>
.form {
    position: relative;
    left: 15px;
    width: 300px;
}
</style>