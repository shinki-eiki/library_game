
<!--  -->
<template>

  <game_depth_one ></game_depth_one>
  <!-- <ws /> -->
</template>

<script >
// import axios  from "axios";
import game_depth_one from './patchouli/game_depth_one.vue'
// import ws from './patchouli/ws.vue'

export default {
    name:"App", // 定义导出名称
    components:{
      game_depth_one,
      // ws,
    },
    data() {
        return {
            // info:
        }
    },
    
}
</script>

<style scoped>

</style>
