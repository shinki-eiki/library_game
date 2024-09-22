
<!-- 符卡展示框的总体设计，测试用例 
父组件需要接受子组件的获取符卡的操作，以及修改某一符卡为新符卡
-->

<template>
  <!-- 目前符卡框大小交由它自身设定 -->

  <div>
    <spellCard @toGetSpell="GetSpellHandler" />
  </div>

  <button @click="addCard">Add Card</button>
  <button @click="popCard">Pop Card</button>

</template>

<script >
import spellCard from "./spellCard.vue"
import { computed } from 'vue'

export default {
  components:{
    spellCard,
  },
  data(){
    return {
      cnt:5,
      spellCard_data: [
        { index: 0, cardName:'金符「Metal Fatigue」' },
        { index: 1, cardName:'木符「Sylphy Horn」' },
        { index: 2, cardName:'水符「Princess Undine」' },
        { index: 3, cardName:'火符「Agni Shine」' },
        { index: 4, cardName:'土符「Lazy Trilithon」' },
      ]
    };
  },
  // created(){
  //   console.log(this.spellCard);
  // },
  provide(){
    return {
      spellCard: computed(() => this.spellCard_data),
      test_text:'hello'
    }
  },
  methods:{
    changeSpell(idx,sp){
      // 接受传过来的数据，将第idx张卡替换成sp
      this.spellCard_data[idx]=sp
    },
    GetSpellHandler(idx){
      // 子组件被点击，将要获取第idx张符卡
      console.log("Will handle to get the spellcard["+idx+"]");
      this.changeSpell(idx, { index: (idx + 1)%5, cardName:'None'})
    }
  }

}
</script>

<style scoped>

</style>
