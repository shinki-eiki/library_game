
<!-- 
  玩家信息展示框的总体设计、测试用例 
  当需要选择某一位玩家时，需要点击然后接受
  包括玩家信息数组，当前玩家索引，手牌数组
-->
<template>

  <div class="entire111">
    <template v-for="(v, idx) in playerList">
      <playerFrame 
        class="playerFrame" 
        :class="idx == current ? 'currentPlayer' : ''" 
        v-bind="v" 
        :elementCount="eleCount(this.handCard[idx])"
        :playerIndex="idx"
        @selectPlayer="select"
         />
    </template>
  </div>
  <!--  v-bind="v"是绑定对象的简单写法，相当于把该对象的键值对作为参数传入 -->

</template>

<script>
import playerFrame from "./playerFrame.vue"

export default {
  components: {
    playerFrame,
  },
  data() {
    return {
    }
  },
  emits: ['selectPlayer'],
  inject: ['playerList','current','handCard'],

  methods: {
    select(idx) {
      // 这里会往上传递
      console.log('Total player frame');
      this.$emit('selectPlayer',idx)
    },

    // 统计手牌的元素分布
    eleCount(es){
      // console.log('Summary of HandCard:',es);
      let res=[0,0,0,0,0]
      for (const e of es) {
        res[e.Property]+=1
      }
      return res
    }
  },

}
</script>

<style scoped>

.entire111{
  /* 不用指定宽度而由子组件确定 */
  background-color: transparent;
  /* background-color: grey; */
  display: flex;
  flex-basis: 80%;
  /* margin:2px ; */
}

.playerFrame {
  margin: 0px 2px;
  /* border-style: solid; */
}

.currentPlayer {
  border: dashed lightcoral;
  /* border-width: 10px; */
}
</style>
