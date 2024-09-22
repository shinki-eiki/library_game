
<!-- 子组件向父组件数据传递，使用$emit -->
<template>
<div>
    <p>Parent</p>
    <!-- 传递常量，可以传递任何类型 -->
    <Child_emit message=20  @aboutName="handleDataFromChild" />
    <p>Data from child: {{ dataFromChild }}</p>

</div>

</template>

<script >
import Child_emit from "./child_emit.vue"
export default {
    components:{
        Child_emit
    },
    data() {
        return {
            dataToChild:"Variable",
            dataFromChild: "",
        }
    },
    methods:{
        // 父组件可以更改传递给子组件的数据
        changeData() {
            this.dataToChild="New variable."
        },
        handleDataFromChild(data){
            console.log("Get data from child."+data);
            this.dataFromChild=data
        }
    }
}
</script>

<style ></style>
