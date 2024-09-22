
<!-- 手牌框的整体，包括玩家头像，帮手按钮，各张手牌-->

<template>

    <!-- 逐个显示的手牌组件 -->
    <div class="entire">

        <!-- 最左边的头像 -->
        <div class="avatar">
            <img class="headImage" :src="Object.keys(player).length === 0?'./headImage/11.jpg':`./headImage/${player.imageIndex}.jpg` " alt="Not found.">

            <!-- <div>clickHand</div>
            <div>Helper</div> -->
        </div>

        <!-- 中间的各张手牌 -->
        <div class="handFrame">
            <div v-for="v, idx in this.handCard" :key="idx">
                <img :src=" `./elementCard/${v.Property}_${v.Index}.jpg` " 
                    @click="this.$emit('clickHand',idx)"
                    class="eleCard"
                    :style="{ 'left': idx * baseWidth + '%' }" 
                    alt="HandCard">
            </div>
        </div>
    </div>

</template>

<script >

export default {
    data() {
        return {
            // 当前手牌的总数
        }
    },

    created(){
        // console.log(this.handCard,this.count);
        // console.log(this.test_text);
    },

    props: {
        handCard: {
            type: Array,
            default() {
                return []
            }
        },
        player:{
            type:Object,
            default(){
                return {

                }
            }
        }
    },

    computed:{
        baseWidth(){
            // 只有当手牌数大于阈值时，才调整宽度以重叠显示
            let res = this.handCard.length < 7 ? 13.5 : 100 / this.handCard.length
            // let res = this.handCard.length<7?150:900/ this.handCard.length
            // console.log(res);
            return res
        }
    },
    emits: [ 'clickHand'],
}
</script>

<style scoped>

.entire{
    width: 96%;
    height: 20%;
    display: flex;
    flex-wrap: nowrap;
    border: 2px;
}

.avatar{
    width: 10%;
    height: 100%;
    /* background-color: antiquewhite; */
    display: flex;
    flex-direction: column;
    justify-content: center;
    border: 2px;
}
.headImage{
    width:98%;
    height: 60%;
    /* object-fill:"fill"; */
    /* text-align:"center"; */
    border: solid 2px black;
}

.handFrame{
    width: 90%;
    height: 100%;
    border: solid 4px;
    /* color: ; */
    position: relative;
    /* overflow: hidden; */
}

.eleCard{
    /* float: left; */
    margin: 0px ;
    width: 13.5%;
    height: 98%;
    position: absolute;
    background-color: transparent;
    /* border: solid; */
    /* 拉伸以填满 */
    object-fit: fill;
}

</style>
