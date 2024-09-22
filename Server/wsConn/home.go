/*
这里只涉及首页和websocket连接的建立，
不涉及具体的通信内容
*/
package wsConn

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "http service address")
	// addr     = flag.String("addr", ":8080", "http service address")
	Home_log *log.Logger
)

func init() {
	// 初始化日志输出目的地
	// file, err := os.OpenFile(
	// 	"Home.log",
	// 	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln("Failed to open home.log:", err)
	// }

	// 特定日志器的生成和设定
	Home_log = log.New(
		// io.MultiWriter(file, os.Stdout), //多个目的地的写法
		// io.Discard, //单个目的地
		os.Stdout,
		"[Home]:",
		log.Lmicroseconds|log.Lshortfile) //log.Ldate|
	Home_log.Println("A new server.")

	// 全局日志器的设定
	// log.SetPrefix("Home:")
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	return
}

/* 返回聊天室的网页,只是单纯的请求和响应，不涉及websocket */
func serveHome(w http.ResponseWriter, r *http.Request) {
	Home_log.Println(r.RemoteAddr, "is logining.")
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	Home_log.Println(r.RemoteAddr, "had logined.")
	http.ServeFile(w, r, "home.html")
}

/* 启动一个房间内的通信 */
func RunHub(hub *Hub) {
	// 开启房间监听进程
	flag.Parse()
	go hub.Run()

	// 然后就是等待用户加入
	// 首页的相应响应
	http.HandleFunc("/", serveHome)
	// 建立websocket所用的函数
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})

	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	// 服务器，启动！
	// err := http.ListenAndServe("127.0.0.1:8080",nil)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
