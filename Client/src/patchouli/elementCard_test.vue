<!-- 符卡展示框的总体设计，测试用例 
父组件需要接受子组件的获取符卡的操作，以及修改某一符卡为新符卡
-->

<template>
  <!-- 目前符卡框大小交由它自身设定 -->
  <!-- style="
      background-image: url('./bg.jpg'); 
      background-size: 100% 100%;
      background-repeat: no-repeat;
      width: 1152px ; 
      height: 648px; " -->
  <div>

    <elementCard @selectElement="GetElementHandler" />
  </div>

  <!-- <button @click="addCard">Add Card</button>
  <button @click="popCard">Pop Card</button> -->

</template>

<script>
// <script setup lang="ts">
import elementCard from "./elementCard.vue"
import { computed } from 'vue'

export default {
  components: {
    elementCard,
  },
  data() {
    return {
      cnt: 5,
      elementCard_data: [
        { index: 0, cardName: '金符「Metal Fatigue」' },
        { index: 1, cardName: '木符「Sylphy Horn」' },
        { index: 2, cardName: '水符「Princess Undine」' },
        { index: 3, cardName: '火符「Agni Shine」' },
        { index: 4, cardName: '土符「Lazy Trilithon」' },
      ]
    };
  },
  // created(){
  //   console.log(this.elementCard);
  // },
  provide() {
    return {
      elementCard: computed(() => this.elementCard_data),
      test_text: 'hello'
    }
  },
  methods: {
    changeElement(idx, sp) {
      // 接受传过来的数据，将第idx张卡替换成sp
      this.elementCard_data[idx] = sp
    },
    GetElementHandler(idx) {
      // 子组件被点击，将要获取第idx张卡
      console.log("Will handle to get the elementCard[" + idx + "]");
      this.changeElement(idx, { index: (idx + 1) % 5, cardName: 'None' })
    }
  },
  // emits:["getElement"]

}
</script>

<style scoped></style>
