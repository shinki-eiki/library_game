
<!-- 中间的状态以及帮手框 -->

<template>
    <div class="entire">
        <div class="each helper" @click="this.$emit('selectHelper')" title="帮手">
            <img src="/helper.jpg" class="image" alt="This is button to get helper.">
        </div>

        <div class="each turn">
            {{ round }}
        </div>

        <!-- :background="url(`/state/${currentState}.jpg`)"> -->
        <div class="each state " 
            @click="clickState"
            :class="this.currentState==1?'canBeClick':''"
            :title="state_text[currentState]" >

            <img :src="`/state/${currentState}.jpg`" 
                class="image"
                alt="No image of the state.">
        </div>
        <!-- status -->
    </div>
</template>

<script>
export default {
    data() {
        return {
            // round:0,
            state_text:[
                "平稳无事：没事。",
                "小恶魔「整理整顿」：在自己的回合内，可以舍弃自己持有的一张属性卡，并从场上取得一张。一回合只能进行一次。",
                "分身：当自己的回合开始，要取得属性卡的时候，改为取得２张。",
                "生病：每当玩家入手符卡的时候，必须追加舍弃一张属性卡。（种类随意）",
                "门卫：不能使用帮手卡。（但是可以入手）",
                "助人为乐：在回合內取得帮手卡的玩家，当回合结束时，从属性卡牌库顶抽一张属性卡入手。",
            ]
        }
    },
    methods:{
        clickState(){
            console.log(this.currentState);
            if (this.currentState!==1) {
                return
            }
            this.$emit('selectState')
        }
    },
    emits: ['selectHelper','selectState'],
    props:{
        currentState: {
            type: Number,
            default: 0,
        },
        round: {
            type: Number,
            default: 0,
        },
    },

}
</script>

<style scoped>
.entire{
    width: 20%;
    background-color: transparent;
    /* border: solid 5px; */
    margin: 0px 5px;
    flex-basis: 20%;
    /* margin: auto; */
}

.each{
    width: 96%;
    text-align:center;
    border: solid 2px;
    /* 可以让元素居中对齐 */
    margin:auto; 
    margin-top: 2%;
    /* border-radius: solid 5px; */
    /* display: table-cell; */
    /* vertical-align:middle; */
}
.image {
    width: 100%;
    height: 100%;
    object-fit: fill;
}

.helper {
    height: 35%;
}

.turn {
    height: 20%;
    font-size: 70px;
    color: violet ;
    border: none;
    /* border-color: black; */
    /* padding: auto; */
    /* font-size: 90%; */
}

.state {
    height: 35%;
}

.canBeClick{
    /* box-shadow: 12px 12px 8px rgba(0, 0, 0, 0.5); */
    border: 5px solid #ccc;
    box-shadow: 5px 5px 4px rgba(0, 0, 0, 0.5);
    /* border-width: 5px; */
    /* border: solid 5px; */
}
</style>
