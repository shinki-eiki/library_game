
// export const tipsText="hello"
export const tipsSet=['Select an element card.','Select a hand card,or cancel.']
// export const tipsText=['Select an element card.','Select a hand card,or cancel.']

// export {tipsText}

// 可用点击和相应的提示文本，用于前端中a开头的操作
export const clickInfo=[
    {code:302,text:""}, // 回合玩家的操作
    {code:64,text:"Select a hand card."},    //选择一张手牌
    {code:16,text:"Select an element card."}, // 选择一张元素
]

// 强制将小于10的数字前面加0作为字符串返回
export function intToStr2(a) {
    let res=String(a)
    // if (a<10) {
    //     res='0'+res
    // }
    return res
}
