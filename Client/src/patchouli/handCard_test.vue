
<!-- 玩家信息展示框的总体设计，测试用例 
当需要选择某一位玩家时，需要点击然后接受
-->

<template>
  
  <div style="width: 900px;height: 200px;background-color: antiquewhite;">
    <handCard />
  </div>
  <button @click="addCard">Add Card</button>
  <button @click="popCard">Pop Card</button>

</template>

<script >
import handCard from "./handCard.vue"
import { computed } from 'vue'
// import handCard from "./handCard.vue"
// import handCard from "./components/handCard.vue"

export default {
  components:{
    handCard,
  },
  data(){
    return {
      cnt:7,
      elementCard: [
        { idx: 0, ele: 0, order: 0 },
        { idx: 1, ele: 2, order: 0 },
        { idx: 2, ele: 1, order: 0 },
        { idx: 3, ele: 4, order: 0 },
        { idx: 4, ele: 4, order: 0 },
        { idx: 5, ele: 3, order: 0 },
        { idx: 6, ele: 1, order: 0 }
      ]
    };
  },
  // created(){
  //   console.log(this.elementCard);
  // },
  provide(){
    return {
      items: computed(() => this.elementCard),
      test_text:'hello'
    }
  },
  methods:{
    addCard() {
      // 添加一张手牌
      this.cnt+=1
      this.elementCard.push({
        idx: this.cnt, ele: this.cnt%5, order: 0
      })
    },
    popCard() {
      // 丢弃一张手牌
      this.cnt -= 1
      this.elementCard.splice(1,1)
    }
  }

}
</script>

<style scoped>

</style>
