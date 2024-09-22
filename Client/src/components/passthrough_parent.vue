
<!-- 透传的写法，直接写class，id，style然后赋值即可 -->
<template>
  <passthrough_child class="s1"/>
      <passthrough_child id="id1"/>
</template>

<script >
import passthrough_child from "./passthrough_child.vue";
export default {
    components:{
        passthrough_child
    }
}
</script>

<style >

</style>
