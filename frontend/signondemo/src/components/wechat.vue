<template>
    <h3 class="title">扫一扫登录</h3>

    <el-tooltip class="item" effect="light" content="返回" placement="right">
        <el-button color="#626aef" @click="props.back()" class="back" :icon="DArrowLeft"></el-button>
    </el-tooltip>
    
    <img :src="props.qrdata.url" class="qrcode"/>
       
</template>
  
<script setup>
    import { DArrowLeft } from  '@element-plus/icons-vue'
    import { checklogin } from '@/api/wechat';
    import { useTokenStore } from '@/stores/token';
    var props = defineProps(['qrdata','back'])
    const check = async()=>{
        checklogin(props.qrdata.ticket).then(res=>{
            const r = res.data
            if (r.data.token!=''){
                useTokenStore().token = r.data.token
                ElNotification({
                    showClose: true,
                    message: "登录成功",
                    type: 'success',
                })
                clearInterval(intervalid)
                window.location.href = "/userhome";
            }
        })
    }
    const intervalid = setInterval(check, 1000);
</script>
  
<style scoped>
.qrcode {
    position: absolute;
    left: 20%;
    top: 20%;
    width: 60%;
    height: 60%;
}
.back {
    position: absolute;
    left: 20%;
    top: 82%;
    width: 60%;
    height: 8%;
}
.title {
    position: absolute;
    font-size: 20px;
    font-weight: bold;
    left: 34%;
    top: 11%;
}
</style>