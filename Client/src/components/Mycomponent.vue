
<!--  -->
<template>
  <div class="singlefile">Here is my component.</div>
</template>

<script >
export default {
    data() {
        return {
            message:'',
        }
    }
}
</script>

<style scoped>
/* scoped makes the style only be effect in this single file. */
.singlefile{
    color: red;
}
</style>
