

<!-- 深度为1的父子组件关系，也就是子组件直接触发交由父组件处理 -->

<template>
    <!-- style="background-image: url('./bg.jpg')" -->

    <body class="entire000">
        <!-- /* 如果一个控件，就直接写，默认似乎会换行，多个控件同行就用div括起来 */ -->

        <!-- 最上方的各个玩家的信息 -->
        <div class=" topFrame">
            <playerFrame_test @selectPlayer="selcectPlayerHandler"></playerFrame_test>
            <!-- <textarea class="messageCSS">{{ messageText }} </textarea> -->
        </div>

        <!-- 中间的三个横向框 -->
        <div class="center">
            <spellCard @selectSpell="selectSpellHandler" />
            <gameState @selectHelper="helperHandler" @clickState="stateHandler" 
                :currentState="currentState" 
                :round="round_data"/>
            <elementCard @selectElement="selectElementHandler" />
        </div>

        <!-- 再留点空隙供提示信息和其他按钮 -->
        <div class="tip_button">
            <div class="tipTextCSS"> {{ tipsText }} </div>
            <!-- <button class="two_button">Cancel</button> -->
            <button class="two_button" @click="cancelHandler">Cancel</button>
            <button class="two_button" @click="endHandler">End</button>
        </div>

        <!-- 最下面的玩家手牌框 -->
        <div class="thisPlayerCSS">
            <handCard 
                class="thisPlayerCSS" 
                :handCard="isGaming?handCard_data[localPlayer]:[]" 
                :player=  "isGaming?playerList_data[localPlayer]:[]"
                @clickHand="selectHandHandler"
            ></handCard>
        </div>

        <!-- 测试用的按钮 -->
        <!-- <button @click="handAppend({ Index: 1 })"> 添加手牌</button> -->
        <!-- <button @click="handPop(1)">删掉第一张</button> -->
        <!-- <input v-model="testInput"> -->
        <!-- <button @click="getDataOfGame">获取游戏数据</button>
        <button @click="connect_ws">连接websocket</button>
        <button @click="testGetEle">测试获取元素</button> -->
        <button @click="connectAndBegin">连接服务器</button> 
        <button @click="gameBegin">开始游戏</button> 
        <!-- <button @click="this.ws_conn.close()">关闭websocket</button>  -->
    </body>

</template>

<script>
import { computed, onMounted } from 'vue'
import axios  from "axios";
import playerFrame_test from './playerFrame_test.vue'
import elementCard from './elementCard.vue'
import spellCard from './spellCard.vue'
import handCard from './handCard.vue'
import gameState from './gameState.vue'
import { tipsSet,clickInfo,intToStr2 } from './tipsText.js'
// import { typeof } from 'os';

export default {
    name: "App", // 定义导出名称
    components: {
        elementCard,
        spellCard,
        playerFrame_test,
        handCard,
        gameState
    },
    mounted(){
        console.log("Function mounted:");
        console.log(tipsSet);
        console.log(this.tipsText);
    },
    data() {
        return {
            // 公共信息（集合类型）
            playerList_data: [
                {
                    playerName: 'reimu',
                    imageIndex: 0,
                    playerIndex: 0,
                    score: 0,
                    imageIndex: 0,
                    hasHelper: false,
                    spellCard:[],
                },
                // {
                //     playerName: 'marisa',
                //     imageIndex: 1,
                //     playerIndex: 1,
                //     score: 0,
                //     imageIndex: 0,
                //     hasHelper: false,
                //     spellCard:[],
                // },
            ],
            spellCard_data: [
                { Name:'01', cardName: '金符「Metal Fatigue」',Score:2 },
                { Name:'05', cardName: '木符「Sylphy Horn」',Score:2 },
                { Name:'10', cardName: '水符「Princess Undine」',Score:2 },
                { Name:'15', cardName: '火符「Agni Shine」',Score:2 },
                { Name:'20', cardName: '土符「Lazy Trilithon」',Score:2 },
            ],
            elementCard_data: [
                { Property:0,Index: 0, cardName: '金符「Metal Fatigue」' },
                { Property:1,Index: 1, cardName: '木符「Sylphy Horn」' },
                { Property:2,Index: 2, cardName: '水符「Princess Undine」' },
                { Property:3,Index: 3, cardName: '火符「Agni Shine」' },
                { Property:4,Index: 4, cardName: '土符「Lazy Trilithon」' },
            ],

            // 应当是一个二元数组
            handCard_data:[
                [
                    { Property:0,Index:1 }, { Property:3,Index:3 },
                    { Property:4,Index:4 }, { Property:2,Index:2 },
                ],
                // [
                //     {Property:0,Index: 0 }, {  Property:2,Index: 2 },
                // ],
                // []
            ],
            
             // 单变量
            messageText: 'Test',  //消息框
            tipsText:'This is the tips.', //提示框
            //当前玩家 
            current_data: 0,  
            localPlayer: -1, //该网页对应的玩家 
            totalPlayer:1, //玩家总数
            valueAsking: 0, //等待点击的询问目标值
            currentState:0,  //当前状态
            round_data: 0,  //回合数
            ws_conn:null, // websocket连接对象
            testInput:"110323",
            clickArea:16, // 可点击区域，是一个位数组
            isAskingVar:false, 
            isGaming:false,
            // 是否正在询问某一个值，用来区分是正在发生事件在等待值，还是要触发事件

        }
    },
    provide() {
        return {
            test_text: 'hello',
            playerList: computed(() => this.playerList_data),
            spellCard: computed(() => this.spellCard_data),
            round: computed(() => this.round_data),
            currentState: computed(() => this.currentState),
            elementCard: computed(() => this.elementCard_data),
            handCard: computed(() => this.handCard_data),
            current: computed(() => this.current_data)
            // localPlayer:0, // 子组件一般是用不到这个的
        }
    },
    methods: {
        connectAndBegin(){ // 一键开始游戏
            this.connect_ws()
            this.getDataOfGame()
            // this.gameBegin()
        },
        beNormalState(){
            this.isAskingVar=false
            this.clickArea=clickInfo[0]['code']
            this.tipsText =clickInfo[0]['text']
        },
        testGetEle(){ // 测试websocket的接受命令的处理
            this.handle_message_ws({data:this.testInput})
        },
        ws_send(mes){ // 通过websocket发送消息的统一函数
            console.log('Websocket Will sned '+mes);
            this.ws_conn.send(mes)
        },
        handle_message_ws(evt){ // 处理websocket传递来的信息

            //#region 消息处理前动作
            // console.log(evt);
            // 认为信息只有一行
            var mes = evt.data.split('\n');
            console.log('Websocket receive '+mes);

            // console.log(typeof(message));
            // const encoder = new TextEncoder(); 
            // const uint8Array = encoder.encode(message)
            // // const uint8Array = new Uint8Array(message);
            // console.log(uint8Array);
            // console.log(String(uint8Array));

            //  检查消息是否有效
            // if (typeof message !== 'string' || message.length < 2) {
            //     console.error('Invalid message format.',typeof(message));
            //     return;
            // }
            //#endregion

            for(let message of mes){
                // 提取操作类型和操作细节
                const operationType = message.substring(0, 1);
                // 注意details从零开始都是有效字符
                const details = message.substring(1);
                console.log(operationType,details);

                // 根据操作类型执行相应操作
                let idx,newCard
                switch (operationType) {
                    // 由前端主动发送的操作
                    case '0': // 玩家获取元素的一般流程
                        // 给当前玩家添加第idx张元素
                        // 再将第idx张元素牌改为target，
                        // 消息格式0123，依次代表操作，目标元素索引，新元素的元素类型，卡图序号
                        idx=Number(details[0])
                        this.handAppend(this.elementCard_data[idx])
                        newCard={Property:details[1],Index:details[2]}
                        this.changeElement(idx,newCard)
                        break;

                    case '1': // 玩家获取符卡的一般流程
                        // 给当前玩家添加第idx张符卡
                        // 再将第idx张元素牌改为新卡牌，
                        // 更改状态
                        // 消息格式 110223 ，依次代表操作1，目标符卡索引1，新符卡名称/序号02，分数2，新状态3
                        idx = Number(details[0])
                        this.getSpell(this.spellCard_data[idx])
                        newCard = { 
                            Name: Number(details.substring(1,3)), 
                            // Index: details[1], 
                            Score: Number(details[3]) 
                        }
                        this.changeSpell(idx,newCard)
                        let newState=Number(details[4])
                        this.changeState(newState)
                        break;
                    
                    // 由后端主动发送的操作
                    case '5': // 结束回合
                        this.turnEnd()
                        break
                    case 'a': // 后端正在处理某个事件，并为此询问某值
                        idx=Number(details)
                        this.isAskingVar=true
                        this.clickArea=clickInfo[idx]['code']
                        this.tipsText =clickInfo[idx]['text']
                        console.log(this.clickArea,this.isAskingVar);
                        break
                    case 'b': // 丢弃某一手牌
                        this.handPop(Number(details))
                        break
                    case 'c': // 设置本地玩家的序号
                        this.localPlayer=Number(details)
                        this.isGaming=true
                        break
                    case 'd': // 抽一张元素牌
                        newCard={Property:details[1],Index:details[2]}
                        this.handAppend(newCard)
                        break
                    default:
                        console.error('Unknown operation type:', operationType);
                }
                console.log("Handle finish.");
            }
        },

        //#region 修改数据的操作，
        // 实际上所有的变化都是数据的变化，vue是支持数据变化时触发动作的，但是我就不搞了
        changeElement(idx, ele) {
            // 接受传过来的数据，将第idx张卡替换成sp
            this.elementCard_data[idx] = ele
        },
        changeSpell(idx, sp) {
            // 接受传过来的数据，将第idx张卡替换成sp
            this.spellCard_data[idx] = sp
        },
        changeState(s){
            this.currentState=s
        },
        handAppend(ele) {
            // 默认给当前玩家添加手牌
            this.newMessage('Get a hand card' + ele)
            this.handCard_data[this.current_data].push(ele)
        },
        handPop(idx) {
            // 默认是当前玩家丢弃手牌
            this.newMessage(`Lose the ${idx} hand card.`)
            this.handCard_data[this.current_data].splice(idx, 1)
        },
        getSpell(sp){
            // 当前玩家获取某一符卡
            // console.log("Get spell"+sp);
            // console.log(this.current_data); //this.playerList_data,
            this.playerList_data[this.current_data].spellCard.push(sp)
            this.playerList_data[this.current_data].score+=sp.Score
        },

        turnEnd(){
            // 当前玩家结束回合，轮到下一位
            this.current_data+=1
            if (this.current_data==this.totalPlayer) {
                this.current_data=0
                this.round_data+=1
                console.log(this.round_data);
            }
        },
        //#endregion
        connect_ws(){
            // 建立websocket连接
            // 创建ws连接并添加相应事件处理
            if (! this.ws_conn) {
                // if (window["WebSocket"]) {
                console.log('When websocket,This is',this);
                let conn = new WebSocket("ws://" + "127.0.0.1:15481" + "/ws"); // 8080

                //连接关闭时的信息提示
                conn.onclose = function (evt) {
                    console.log("Connection closed.");
                };

                this.ws_conn = conn
                conn.binaryType = 'arraybuffer';
                // 接受信息时的处理
                conn.onmessage = this.handle_message_ws

            } 
            else { //不支持时的操作
                // var item = document.createElement("div");
                // item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
                console.log('Your browser does not support WebSockets.');
                
            }
        },
        gameBegin(){
            axios.get('http://127.0.0.1:15481/begin')
        },
        getDataOfGame(){
            // 获取游戏的相关数据
            let tp=this
            axios.get('http://127.0.0.1:15481/begin/1')
            .then(function (response) {
                // console.log(response);
                console.log('Data of Game:',response.data);
                let data = response.data
                console.log(this);
                console.log(tp);
                tp.elementCard_data = data.element
                tp.spellCard_data = data.spell

                // 玩家信息数据的处理
                let idx=0
                tp.handCard_data=[]
                tp.playerList_data=[]
                tp.totalPlayer=data.players.length
                for (const p of data.players) {
                    console.log(p.Hand);
                    tp.handCard_data.push(p.Hand)
                    tp.playerList_data.push({
                        playerIndex : idx,
                        playerName : p.Name,
                        imageIndex : idx,
                        score : 0,
                        hasHelper : false,
                        spellCard:[],
                    }),
                    idx++
                }
                
                // 其他单变量的设置
                tp.beNormalState()
            }).catch(function (error) {
                console.log(error);
            })
            return
        },

        // 点击某一按钮后的修饰函数。
        // 接受一个值以及一个函数，判断该玩家是否当前玩家，再判断当前区域是否可以点击
        afterClick(idx,fun){
            if (this.localPlayer!=this.current_data) {
                return
            }
            if ( (1<<idx)&this.clickArea==0) {
                return
            }
            return fun()
        },
        // 上面的非修饰函数形式,area代表所点击的区域，返回事件是否不处理的布尔值
        beforeHandle(area){
            if (this.localPlayer!=this.current_data) {
                return true
            }
            console.log(this.clickArea,1<<area,(1<<area)&this.clickArea);
            if ( ((1<<area)&this.clickArea)===0) {
                return true
            }
            console.log('Click area'+area+' successfully.');
            return false
        },

        // 给消息框发送消息
        newMessage(s){
            this.messageText+='\n'+s
            // this.messageText.push(s)
        },
        
        // 发送所询问的索引消息，并恢复正常
        returnAskingVar(idx){
            let mes="a"+intToStr2(idx)
            this.ws_send(mes)
            console.log('Is asking?',this.isAskingVar);
            this.isAskingVar=false
            this.beNormalState()
            console.log('Is asking?',this.isAskingVar);
        },
        // 会因点击触发的子组件包括：玩家，元素，符卡，帮手，状态，手牌，自己的帮手
        selcectPlayerHandler(idx) { 
            if (this.beforeHandle(0)) {return}
            this.returnAskingVar(idx)
        },
        // 中间框从左往右
        selectSpellHandler(idx) {
            if (this.beforeHandle(1)) {return}
            if (this.isAskingVar) {
                return this.returnAskingVar(idx)
            }
            this.newMessage("You select spellcard" + idx);
            this.ws_send("1" + idx);
        },
        helperHandler() {
            if (this.beforeHandle(2)) {return}
            this.newMessage("You select helper");
        },
        stateHandler() {
            console.log("Click State");
            if (this.beforeHandle(3)) {return}
            this.newMessage("You click state.");
        },
        selectElementHandler(idx) {
            if (this.beforeHandle(4)) {return}
            if (this.isAskingVar) {
                return this.returnAskingVar(idx)
            }
            // 元素组件被点击，表示将要获取第idx张卡
            this.newMessage("Will handle to get the elementCard[" + idx + "]");
            this.ws_send('0'+idx)
        },
        endHandler() {
            console.log("Click End");
            if (this.beforeHandle(8)) {return}
            this.ws_send('5')
            this.newMessage("End you turn.");
        },
        // 按钮以及手牌的处理
        selectHandHandler(idx) {
            console.log("Click hand " +idx);
            if (this.beforeHandle(6)) {return}
            if (this.isAskingVar) {
                return this.returnAskingVar(idx)
            }
            this.newMessage("You choose handcard"+idx);
        },
        confirmHandler() {
            if (this.beforeHandle(11)) {return}
            this.newMessage("You confirm.");
        },
        cancelHandler() {
            if (this.beforeHandle(7)) {return}
            this.newMessage("You cancel.");
        },
        
        useHelperHandler() {
            if (this.beforeHandle(5)) {return}
            this.newMessage("You will use helper state.");
        },
        
        
    },
}
</script>

<style scoped>

*{
    --totalWidth:1200px;
}

.entire000{
    width: var(--totalWidth);
    /* background: url('./pat1.jpg') no-repeat;
    background-size:cover; */
    /* background-image: url('./pat0.jpg'); */
    background-color: beige;
    /* z-index: -2; */
    padding: 3px;
    border: dashed 3px;
    border-bottom: 0px;

}
.topFrame{
    width: var(--totalWidth);
    height: 120px;
    display: flex;
    margin-bottom: 8px;
    /* background-color: antiquewhite; */
    /* backgr */
}

.messageCSS {
    width: 20%;
    height: 95%;
    color: lightgreen;
    background-color: darkcyan;
}

/* 中间的三个框的组合框 */
.center {
    width: var(--totalWidth);
    height: 425px; 
    background-color:transparent;
    /* background-color:blanchedalmond;*/

    display: flex;
    /* 下面语句代表行内溢出时不换行 */
    margin-bottom: 6px;
    margin-right: 6px;
}

.tip_button{
    width: var(--totalWidth);
    height: 50px;
    background-color: transparent;
    /* background-color: lightgrey; */
    display: flex;
    justify-content: flex-end;
    align-items: center;
}
.tipTextCSS{
    text-align: center;
    flex-basis: 80%;
    color:violet;
}
.two_button{
    width:8%;
    height:80%;
}

.thisPlayerCSS{
    width: var(--totalWidth);
    height: 200px;
}


</style>
