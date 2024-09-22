
<!-- 
    符卡框整体
可以点击图片框触发事件
只有当前玩家点击，并且当回合未获得过符卡有效，需要上级判断

emits：点击某一符卡，代表当前玩家想要获取某一符卡
由于只是展示用，因此只需要接受符卡的图片索引，展示是哪一张即可，消耗和分数都写在图片上
注意符卡的模型里面是没有索引这种说法的，以后可能会加上，
以及会有重复的符卡出现

后端模型定义如下
type SpellCard struct {
	Name  string
	Cost  [5]int
	Score int
}
-->

<template>

    <div class="entire">
        <!-- 需要双重嵌套，不然不是设想的表现 ，template相当于一个虚拟的父组件，实际操作发现是可行的-->
        <template v-for="v, idx in this.spellCard" :key="idx">
            <img :src="`./spellCard/${v.Name}.jpg`" 
                @click="toGetSpell(idx)" 
                class="spell"
                :style="{ 'left': place[0][idx]+'%' , 'top':place[1][idx]+'%' }" 
                alt="No spell_card">
            <!-- alt可以改成显示符卡的name -->
            <!-- <p>{{ v.cardName }}</p> -->
        </template>

    </div>

    <!-- <div>{{ test_text }}</div> -->
</template>

<script >
export default {
    data() {
        return {
            // 定义五张牌各自的位置，这里用相对位置做
            place:[
                [15,55,0,35,70], // x
                [2,2,52,52,52]  // y 
            ]
        }
    },

    // created(){
    //     console.log(this.spellCard);
    // },

    inject:[ 'spellCard','test_text'],
    emits: [ 'selectSpell'],

    methods:{
        toGetSpell(which){
            // 点击某一符卡，将进行获取之的请求发送
            // console.log(which);
            console.log("Click spellcard no."+which);
            this.$emit("selectSpell",which)
        }
    }

}
</script>

<style scoped>
.entire{
    width: 40%;   
    background-color: transparent;
    border: solid;
    flex-basis: 40%;
    position: relative;
}

.spell{
    /* 高度占40%，宽度占30%，所以总框是长500宽500px */
    /* height: 200px;
    width: 150px;  */
    height: 46%;
    width: 30%; 
    position: absolute;
    /* background: antiquewhite; */
}
</style>
