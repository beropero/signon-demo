<template>
    <el-form
    label-position="top"
    label-width="auto"
    :model="form"
    class="form"
    >
    <el-form-item label="账号">
      <el-input v-model="form.account" clearable placeholder="请输入邮箱/手机号" maxlength="30">
        <template #append>
            <el-tooltip class="item" effect="dark" content="发送验证码" placement="bottom-start">
                <el-button type="primary" :icon="Promotion" @click="sendCode()"/>
            </el-tooltip>
           
        </template>
      </el-input>
    </el-form-item>
    <el-form-item label="验证码">
      <el-input v-model="form.code" clearable placeholder="请输入验证码" maxlength="6"/>
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
    import {Promotion} from '@element-plus/icons-vue'
    import { emailSendCode } from '@/api/email';
    import { emailSignin } from '@/api/email';
    import { mobileSendCode } from '@/api/mobile';
    import { mobileSignin } from '@/api/mobile';
    import { useTokenStore } from '@/stores/token';
    const formRef = ref<FormInstance>()
    const form = reactive({
        account: '',
        code: '',
    })
    function sendCode() {
        if(form.account.includes('@')){
            emailSendCode(form).then((res) => {
                const r = res.data
                if (r.code == 0) {
                    ElNotification({
                        showClose: true,
                        message: "发送成功",
                        type: 'success',
                    })
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
            mobileSendCode(form).then((res) => {
                const r = res.data
                if (r.code == 0) {
                    ElNotification({
                        showClose: true,
                        message: "发送成功",
                        type: 'success',
                    })
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