
<!-- 父子数据传递，使用props -->
<template>
<div>
    <p>Parent</p>
    <!-- 传递常量，可以传递任何类型 -->
    <Child message=20 />

    <!-- 传递变量，可以传递任何类型 -->
    <Child :message="dataToChild" />
    <!-- // 更改传递给子组件的数据 -->
    <button @click="changeData">Change data to child.</button>
</div>

</template>

<script >
import Child from "./child.vue"
export default {
    components:{
        Child
    },
    data() {
        return {
            dataToChild:"Variable"
        }
    },
    methods:{
        // 父组件可以更改传递给子组件的数据
        changeData() {
            this.dataToChild="New variable."
        }
    }
}
</script>

<style ></style>
