<template>
    <div class="container">
        <div class="context">
            <div class="left">
                <p>Signon Demo</p>
            </div>
            <div class="right">
                <div class="form">
                    <SignUp v-if="show.official"/>
                    <Wechat :qrdata="qrdata" :back="back" v-if="show.wechat"/>
                </div>
                <div class="other">
                    <p>其他登录方式</p>
                    <hr width="80%" color=black size=0.1/>
                    <button id="myButton" @click="wechatlogin()">
                        <img src="../assets/wechat.svg" alt="WeChat Icon" style="position: relative;top: -15%;left: -40%; width: 30px; height: 30px;">
                    </button>
                </div> 
            </div>
        </div>
    </div>
</template>

<script setup>
    import SignUp from "../components/SignUp.vue"
    import Wechat from "@/components/wechat.vue";
    import { reactive, ref } from 'vue'
    import { getQrcode } from "@/api/wechat";
    var show = reactive({
        official: true,
        wechat: false
    })
    const qrdata = reactive({
        ticket: '',
        url: '',
    })
    function back(){
        show.official = true
        show.wechat = false
    }
    function wechatlogin(){
        getQrcode().then(res => {
            const r = res.data
            qrdata.ticket = r.data.ticket
            qrdata.url = r.data.url
            show.official = false
            show.wechat = true
        }) 
    }
</script>

<style scoped>
div {
    position: relative;
}
.container {
    background-image: url("../assets/background.jpg");
    background-repeat: no-repeat;
    background-size: cover;
    display: flex;
    align-items: center;
    justify-content: center;
    
    width: 100%;
    height: 100%;
    
}
.context { 
    display: flex;
    background-image: url("../assets/background.jpg");
    background-repeat: no-repeat;
    background-size: cover;
    width: 1000px;
    height: 600px;
    border-radius: 30px;
    box-sizing: border-box;
    box-shadow: 0px 4px 10px 0px rgba(0, 0, 0, 0.3);
}
.left {
    left: 200px;
    top: 180px;
    width: 200px;
    height: 200px;
    font-size:50px
    
}
.right {
    position: absolute;
    left: 600px;
    top: 50px;
    width: 350px;
    height: 500px;
    background: rgba(255, 255, 255, 0.5);
    border-radius: 30px;
    box-sizing: border-box;
    box-shadow: 0px 4px 10px 0px rgba(0, 0, 0, 0.3);
}
.form {
    left: 10px;
    top: 10px;
    width: 330px;
    height: 350px;
}
.other {
    font-size:12px;
    display: flex;
    flex-direction: column;
    align-items: center;
    left: 10px;
    top: 10px;
    width: 330px;
    height: 130px;
}
#myButton {
    margin: 10px;
    border-radius: 100%; 
    width: 25px;
    height: 25px;
    border: none; 
    outline: none; 
    cursor: pointer; 
}

</style>