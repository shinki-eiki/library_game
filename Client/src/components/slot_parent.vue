
<!--  -->
<template>
    <p>Parent</p>
  <slot_child>
    <p>Slot to child</p>
    <a href="">Just a href.</a>
  </slot_child>
</template>

<script >
import slot_child from "./slot_child.vue"
export default {
    data() {
        return {
            
        }
    },
    components:{
        slot_child
    }
}
</script>

<style >

</style>
