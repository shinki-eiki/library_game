// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wsConn

import (
	"fmt"
	"log"
	"os"
)

var (
	// 记录房间信息的日记器
	Hub_log *log.Logger
)

// 主要是日志输出的初始化
func init() {
	// 初始化hub的日志输出目的地
	// file, err := os.OpenFile(
	// 	"Hub.log",
	// 	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln("Failed to open hub.log:", err)
	// }

	// 特定日志器的生成和设定
	Hub_log = log.New(
		// io.MultiWriter(file, os.Stdout), //多个目的地的写法
		os.Stdout,
		"[Hub]:",
		log.Lmicroseconds|log.Lshortfile) //log.Ldate|
	Hub_log.Println("A new hub.")
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
// 因为是实例，所以很多都是以简单为主的写法
// 虽然有些是仅读或者仅写的管道，但是其实定义成双向管道无所谓的
// 以及这里的用户记录的字典实际上是当集合用的，毕竟没有集合这种东西嘛
type Hub struct {
	// Registered clients.
	clients    map[*Client]bool
	clientList []*Client
	// Inbound messages from the clients.
	// 任一客户端发送的消息。
	broadcast chan []byte
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
	// 游戏相关的管道
	toGame   chan []byte
	fromGame chan []byte
	// Count of clients in this hub.
	count int
}

// 单纯地产生一个Hub对象
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		toGame:     make(chan []byte),
		fromGame:   make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		clientList: make([]*Client, 0),
		count:      0,
	}
}

// 获取与游戏通信相关的管道，依次为发送到游戏和接受游戏的信息的管道
func (h *Hub) GetChannelOfGame() (toGame chan []byte, fromGame chan []byte) {
	return h.toGame, h.fromGame
}

// 一名用户退出时，进行的移除操作
func (h *Hub) delete(client *Client) {
	if _, ok := h.clients[client]; ok {
		for idx, v := range h.clientList {
			if v == client {
				h.clientList = append(h.clientList[:idx], h.clientList[idx+1:]...)
				break
			}
		}
		delete(h.clients, client)
		close(client.send)
		h.count--
		Hub_log.Println(client, "unregistered")
	}
}

// 不断处理用户加入，退出，发送消息的事件
func (h *Hub) Run() {
	for {
		select { //没有default语句，因此将阻塞直到有管道可执行

		// 用户加入时，用户列表只加入不删除，注意一下
		case client := <-h.register:
			h.clients[client] = true
			h.clientList = append(h.clientList, client)
			h.count++
			Hub_log.Println(client, "client registered")

		// 用户退出
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.delete(client)
			}

		// 处理用户发送过来的信息
		case message := <-h.broadcast:
			// 打印一下接受到的内容
			// 用日记记录发送者，然后逐字符打印内容
			Hub_log.Printf("Receive message from client:%v\n", message)
			// for i, v := range message {
			// 	fmt.Println(i, v, string(v))
			// 	// Hub_log.Println(i, string(v))
			// }
			// fmt.Println(message)

			// 交由Game处理消息，处理结果不在这里处理，
			// 而是通过fromGame另一个管道统一发送
			fmt.Printf("Will send to %v\n", h.toGame)
			h.toGame <- message
		}
	}
}

// 不断接受来自game的消息，然后转发到各个client
func (h *Hub) MessageFromGameHandle() {
	var mes []byte
	for {
		mes = <-h.fromGame
		Hub_log.Printf("Receive message from Game:%v\n", mes)
		if len(mes) == 0 { // 发送空消息代表终止管道
			// close(h.fromGame)
			break
		}

		target := h.clientList //默认情况下向所有用户发送
		// 当开头为-时，设置为单独向某一用户发送（即mes[1]）
		if mes[0] == '-' {
			target = []*Client{h.clientList[mes[1]]}
			mes = mes[2:]
		}

		// 给目标用户发送处理结果，包括发送者
		fmt.Println("Sending to ", target)
		for _, client := range target {
			// _, ok := <-client.send
			// if !ok {
			// 	close(client.send)
			// 	delete(h.clients, client)
			// 	Hub_log.Println(client, "disconnect")
			// 	continue
			// }

			select {
			case client.send <- mes:
			default:
				// 如果上面的管道走不通，则执行default，认为用户连接已经不存在了
				close(client.send)
				delete(h.clients, client)
				Hub_log.Println(client, "disconnect")
			}
		}

	}
}
