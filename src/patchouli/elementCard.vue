<!-- 
    元素卡框整体
可以点击图片框触发事件
行动阶段的首个动作就是获取一张元素牌

emits：点击某一符卡，代表当前玩家想要获取某一符卡
由于只是展示用，因此只需要接受符卡的图片索引，展示是哪一张即可，卡牌信息都写在图片上
注意符卡的模型里面是没有索引这种说法的，以后可能会加上，

后端模型定义如下
type Element struct {
	Property int //Element type
	Index    int //Image index
}Element
-->
<template>

    <div class="entire">
        <!-- 需要双重嵌套，不然不是设想的表现 ，template相当于一个虚拟的父组件，实际操作发现是可行的-->
        <template v-for="v, idx in this.elementCard" :key="v.idx">

            <img 
                :src="`./elementCard/${v.Property}_${v.Index}.jpg`" 
                @click="select(idx)" 
                class="Element"
                :style="{ 'left': place[0][idx] + '%', 'top': place[1][idx] + '%' }" 
                alt="No Element_card">

        </template>

    </div>

</template>

<script>
export default {
    data() {
        return {
            // 定义五张牌各自的位置，这里用相对位置做
            place: [
                [15, 55, 0, 35, 70], // x
                [2, 2, 52, 52, 52]  // y 
            ]
        }
    },

    created() {
        // console.log(this.place);
    },

    inject: ['elementCard', 'test_text'],
    emits: ['selectElement'],

    methods: {
        select(which) {
            // 点击某一符卡，将进行获取之的请求发送
            console.log("Click Elementcard no." + which);
            this.$emit("selectElement", which)
        }
    }

}
</script>

<style scoped>

/* 整个框的样式 */
.entire {
    width: 40%;
    background-color: transparent;
    border: solid;
    flex-basis: 40%;
    position: relative;
    /* 由于子组件是固定位置所以必须设置relative */
}

/* 卡片的样式 */
.Element {
    /* margin: 0px 10px; */
    /* 高度占40%，宽度占30%，所以总框是长500宽500px */
    height: 46%;
    width: 30%; 
    position: absolute;
    /* background: antiquewhite; */
}
</style>
