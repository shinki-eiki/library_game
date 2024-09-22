
<!-- 接受透传的子节点，注意必须只能有一个根节点，然后对该根节点及其子节点都有影响 -->
<template>
  <div>
    root_div
    <p >p1</p>
    <p class="s2">p2</p>
  </div>
  <!-- <div>div2</div> -->
</template>

<script >
export default {
    data() {
        return {
            
        }
    }
}
</script>

<style scoped >
.s1{
    color: red;
}
.s2{
    color: blue;
}
#id1{
    color: green;
}
</style>
