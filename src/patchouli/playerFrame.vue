
<!-- 单个玩家的信息显示框体，
分四个div：头像，元素计数，名称，其他信息-->

<!-- 属性：imageIndex头像索引,
 名称playerName，5种元素的计数elementCount， -->
<template>
    <!-- Entire frame: -->
    <!-- 给整个区域添加一个点击响应时间，
        TODO : 注意需要添加一个状态标记以确定是否可用 -->

    <div @click="select(playerIndex)" class="entire">

        <div class="imageFrame">
            <img :src="`/headImage/${imageIndex}.jpg`" class="avatar" alt="">
        </div>

        <div class="elementFrame">
            <div v-for="(val, idx) in elementCount" :key="idx">
                <p :class=" `elementNumber element${idx}`">{{ val }} </p>
            </div>
        </div>

        <div class="nameFrame">{{ playerName }}</div>
        <div class="infoFrame" :title="title_text">{{ info_text }}</div>

    </div>

</template>

<script >
export default {
    
    data() {
        return {
            //spellCard: [],
        };
    },
    methods:{
        select(){
            console.log("select one player");
            // alert("You select player "+this.playerIndex)
            this.$emit('selectPlayer',this.playerIndex)
        },
        mounted() {
            console.log(this.imageIndex);
        }
    },
    computed:{
        info_text(){
            // 包括得分，帮手、已有符卡等的信息文本
            let helperLabel=this.hasHelper? "H" :"N"
            let cardCnt=this.spellCard.length
            // console.log(helperLabel,cardCnt);
            return `${this.score} ${cardCnt} ${helperLabel}`
            // return "infomation"
        },

        title_text(){ // 头像框的详细介绍
            let s = `${this.playerName} has ${this.spellCard.length} spell cards,and\
 his score is ${this.score}.`
            let helperLabel = this.hasHelper ? "Has a helper." : "Has no helper."
            return s+helperLabel
        }
    },
    emits: ['selectPlayer'],
    
    props:{
        playerIndex: {
            type: Number,
            default: 0,
        },
        playerName:{
            type:String,
            default:'NoName'
        },
        imageIndex:{
            type:Number,
            default:0,
        },
        score:{
            type: Number,
            default: 0,
        },
        hasHelper: {
            type: Boolean,
            default: false,
        },
        // 五种元素的统计数组
        elementCount: {
            type: Array,
            default(){
                return [0, 0, 0, 0, 0]
            },
        },
        spellCard: {
            type: Array,
            default() {
                return []
            },
        },
    }

}
</script>

<style scoped>

*{
    margin: 0px;
    border-width: 0px;
    /* border: solid; */
}

:root {
    --test: 30px;
    /* --entire_length: calc(var(--base_length)*5px); */
}

.entire {
    --base_length: 22px;
    --base_font_size: 17px;
    /* width: calc(var(--base_length) * 3.6); */
    height: calc(var(--base_length) * 5);
    position: relative;
    flex-basis: 10%;
    background-color: transparent;
    border:solid 3px black;
}

.imageFrame {
    width: calc(var(--base_length) * 3);
    height: calc(var(--base_length) * 3);
    background-color: grey;
}

.avatar {
    width: 100%; 
    object-fit: fill;
}

.nameFrame {
    position: absolute;
    top: calc(var(--base_length) * 3);
    width: calc(var(--base_length) * 3);
    height: var(--base_length);
    /* color: #CCFFFF; */
    background-color: #99CCCC;
    font-size: var(--base_font_size);
    text-align: center;
}

.elementFrame {
    position: absolute;
    top: 0px ;
    left: calc(var(--base_length) * 3);
    width: var(--base_length) ;
    height: calc(var(--base_length) * 5);
    /* color: paleturquoise; */
    float: right;
    background-color: whitesmoke;
    /* background-color: #CCCCCC; */
    text-align: center;

}

.infoFrame {
    position: absolute;
    top: calc(var(--base_length) * 4);
    width: calc(var(--base_length) * 3);
    height: var(--base_length);
    /* color: palegreen; */
    font-size: var(--base_font_size);
    /* background-color: palegreen; */
    background-color: #CCFFCC;
    text-align: center;
}

p {
    width: var(--base_length * 1px);
    height: var(--base_length * 1px);
}

/* 元素数字字体大小以及5种颜色 */
.elementNumber{
    font-size: var(--base_font_size);
}

.element0 {
    color: gray;
}

.element1 {
    color: green;
}

.element2 {
    color: blue;
}

.element3 {
    color: red;
}

.element4 {
    color: goldenrod;
    /* color: wheat; */
}

</style>
