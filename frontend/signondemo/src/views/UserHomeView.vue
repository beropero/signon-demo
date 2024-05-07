<template>

</template>

<script setup>
    import { getuserinfo } from '@/api/user';
    function init() {
        getuserinfo().then(res=>{
            console.log(res.data)
        })
    }
    init()
</script>