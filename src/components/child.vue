
<!--  -->
<template>
    <div>
        Child
        <br>
        {{ message }}
        <br>
        Name:{{ name }}
        <br>
        Hobby:{{ hobby }}
    </div>
</template>

<script >
export default {
    data() {
        return {

        }
    },
    /* 子组件使用props接受父组件的数据
    子组件只能读取而不能修改数据，而父组件则可以
     */
    props:{
        message:{
            // 指定接受数据类型
            type:[Number,String]
        },
        name:{
            type:String,
            // 未接收到值时使用默认值
            default:"Noname",
            // 是否必须传递的值？
            // required:True,
        },
        hobby:{
            type:Array,
            // 对于数组以及对象，默认值需要使用工厂函数
            default() {
                return ['No habits.']
            }
        }
    }
}
</script>

<style ></style>
