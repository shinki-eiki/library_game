
<!--  -->
<template>

    <input type="text" v-model.lazy="target"> 
</template>

<script >
export default {
    data() {
        return {
            target:'',
        }
    },
    watch:{
        target(nv,ov){
            this.$emit('changeTarget',nv)
        }
    }
}
</script>

<style ></style>
