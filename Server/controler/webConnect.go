package controler

import (
	"fmt"
	"net/http"
	"patchouli_lib/model"
	"patchouli_lib/wsConn"

	"math/rand"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	HUB *wsConn.Hub
	// GAME
)

// 测试用的打招呼函数
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

// 获取游戏数据
func GetDataOfGame(c *gin.Context) {
	// 单纯地返回整个游戏数据的json

	// g := model.NewGame([]string{"reimu", "marisa"})
	// idx := c.Param("index")
	// m := make(map[string]bool, 0)
	// m[idx] = true
	// 可以连续对c的内容进行操作，函数结束时才发送响应消息
	// c.String(http.StatusOK, idx+"\n")
	// Bind the JSON data from the request body to the User struct
	// if err := c.ShouldBindJSON(&GAME); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, GAME)
	// time.Sleep(time.Second * 2)
	// GAME.Begin()
}

// 建立websocket通信
func serveWS(c *gin.Context) {
	w := c.Writer
	r := c.Request
	wsConn.ServeWs(HUB, w, r)
}

func gameBegin(c *gin.Context) {
	GAME.Begin()
	c.String(http.StatusOK, "Game begin!")
}

// 设置gin和websocket的连接
func AllConnect() {
	fmt.Println("Game Begin.")
	rand.Seed(42)

	// 定义websocket以及将其与Game连接
	HUB = wsConn.NewHub()
	from, to := HUB.GetChannelOfGame()
	go HUB.Run()
	go HUB.MessageFromGameHandle()
	GAME = model.NewGame([]string{"Reimu", "Marisa"}, to, from) //, "Alice"  ,, "Marisa", "Alice"

	// gin框架路由的初始化
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", hello)
	r.GET("/begin", gameBegin)
	r.GET("/begin/:index", GetDataOfGame)
	r.GET("/ws", serveWS)
	err := r.Run("127.0.0.1:15481")
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
