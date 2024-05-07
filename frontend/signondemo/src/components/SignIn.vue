<template>
    <el-tabs 
        stretch
        v-model="activeName" 
        class="demo-tabs" 
        @tab-click="handleClick"
    >
      <el-tab-pane label="账密登录" name="first">
        <bypass/>
      </el-tab-pane>
      <el-tab-pane label="验证码登录" name="second">
        <bycode/>
      </el-tab-pane>
    </el-tabs>
</template>
<script lang="ts" setup>
  import { ref } from 'vue'
  import type { TabsPaneContext } from 'element-plus'
  import bypass from "@/components/form/login/pass.vue"
  import bycode from "@/components/form/login/code.vue"
  const activeName = ref('first')
  
  const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
  }
</script>
  
<style scoped>
.demo-tabs {
    width: 100%;
    font-size: 32px;
    font-weight: 600;
}
</style>
  