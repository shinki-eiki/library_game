
<!-- 演示怎样通过$emit发送数据给父组件 -->
<template>
    <div>
        Child=====================
        <br>
        {{ message }}
        <br>
        Name:{{ name }}
        <br>
            <button @click="sendToFather">Send data to parent.</button>
    </div>
</template>

<script >
export default {
    data() {
        return {

        }
    },
    props:{
        message:{
            type:[Number,String]
        },
        name:{
            type:String,
            default:"Noname",
        },
    },
    methods:{
        sendToFather(){
            this.$emit("aboutName",'new name.')
        }
    }
}
</script>

<style ></style>
