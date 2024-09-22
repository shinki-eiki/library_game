
<!-- 父组件接受子组件搜索框的变化 -->
<template>
<div>
    <p>Parent</p>
    <!-- 传递常量，可以传递任何类型 -->
    <searchFrame @changeTarget="handleDataFromChild" />
    <p>You are searching for {{ dataFromChild }}</p>

</div>
</template>

<script >
import searchFrame from "./searchFrame.vue"
export default {
    components:{
        searchFrame
    },
    data() {
        return {
            dataFromChild: "",
        }
    },
    methods:{
        handleDataFromChild(data){
            console.log("Get data from child."+data);
            this.dataFromChild=data
        }
    }
}
</script>

<style ></style>
