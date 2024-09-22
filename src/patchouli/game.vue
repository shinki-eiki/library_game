
<!--  -->
<template>
    <playerFrame_test></playerFrame_test>
    <br>
    <!-- <template style="float: left;"> -->
        <div>Spell_Card</div>
        <spellCard_test></spellCard_test>
    <elementCard_test></elementCard_test>
    <handCard_test></handCard_test>
</template>

    <script>
import elementCard_test from './elementCard_test.vue'
import spellCard_test from './spellCard_test.vue'
import playerFrame_test from './playerFrame_test.vue'
import handCard_test from './handCard_test.vue'

export default {
    name:"App", // 定义导出名称
    components:{
        elementCard_test,
        spellCard_test,
        playerFrame_test,
        handCard_test,
    },
    data() {
        return {
            // info:
        }
    },
    
}
</script>

    <style>

    </style>
