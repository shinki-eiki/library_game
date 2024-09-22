<!--  -->
<template>
    <div>Message:</div>
    <div>{{ this.text }}</div>
    <!-- <br> -->
    <div>-----------------------------------------------------</div>
    <form @submit.prevent="sendToWs">
        <input type="text" v-model="formText" size="64" autofocus />
        <input type="submit"  />
    </form>

</template>

<script>
export default {
    data() {
        return {
            conn: undefined,
            text: "",
            formText: ''
        }
    },
    mounted() {
        if (window["WebSocket"]) {
            this.conn = new WebSocket("ws://localhost:8080/ws");
            let instance = this
            //连接关闭时的信息提示
            // 注意这个回调函数的this是websocket而不是组件实例
            // 默认是继承当前所处对象也就是ws的
            // 所以外部声明一个变量指向实例
            this.conn.onclose = function (evt) {
                instance.text = "Connection closed."
            };

            // 接受信息时的处理
            this.conn.onmessage = function (evt) {
                var messages = evt.data //.split('\n');
                console.log('Receive :',evt.data);
                instance.newMessage(messages)
                // this.text = messages
            };
            this.text="Built connection."
        }
    },
    methods: {
        sendToWs() {
            if (!this.conn) {
                return false;
            }
            if (!this.formText) {
                return false;
            }
            console.log(this.formText);
            this.conn.send(this.formText);
            this.formText = "";
            return false;
        },

        newMessage(mes){
            console.log("Will change to",mes);
            this.text=mes
        }
    }
}
</script>

<style></style>
