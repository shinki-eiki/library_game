/* 单独与一个用户对接的websocket */

package wsConn

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	// "strconv"
	"time"

	"github.com/gorilla/websocket"
)

// 连接过程的常量
const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// 格式符号定义
var (
	newline     = []byte{'\n'}
	space       = []byte{' '}
	clientNames = []string{"reimu", "marisa", "sanae", "alice", "youmu", "sakuya"}
)

// 产生websocket的工厂
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// 房间与前端ws之间的中间件
// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	// 定义时是用256个byte的通道的，也就是有长度的
	send chan []byte

	// Index of this client.
	index int

	// Name of this client.
	name string
}

func (c *Client) String() string {
	// idx:=strconv.Itoa(c.index)
	return fmt.Sprintf("[%v]%v", c.index, c.name)
}

// readPump pumps messages from the websocket connec c.index,c.to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
// 读取前端传来的信息，通过broadcast管道交由Hub处理
func (c *Client) readPump() {
	// 保证最后关闭连接，并且通知房间已退出
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	// websocket的参数设置
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// 读取网页传来的信息并处理，然后传给广播管道
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// 先将字符串中原来的换行变成空格，再除去前后的空白，一般情况下是不会改变原字符串的
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// 添加上发送者的索引，然后发送到hub
		res := bytes.NewBuffer([]byte{})
		// res := bytes.NewBuffer([]byte{byte(c.index) + 48})
		// res := bytes.NewBuffer([]byte{byte(c.index) + 48, ':', ' '})
		fmt.Println("[Client] Get message from front:user", c.index)
		res.Write(message)
		c.hub.broadcast <- res.Bytes()

		// head:=[]byte{byte(c.index)}
		// message = []byte(fmt.Sprintf("%v:%v", c.index, message))
		// message=append(head,message[:] )
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
// 由send管道接收到从hub发送来的信息，然后传递到前端
func (c *Client) writePump() {
	// 定时向管道发送信息，以查看连接状况是否正常
	ticker := time.NewTicker(pingPeriod)
	defer func() { // 最后保证管道关闭以及定时器关闭
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		// 接收到从hub发送来的信息
		case message, ok := <-c.send: // 看起来这句话只是测试通道是否通畅
			// 设置发送时间限制，并处理hub已经关闭管道的情况
			fmt.Println("[Client] Receive from Hub:", message)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait)) //相当于当前时间往后10s
			if !ok {                                           // ok是用来判断管道是否已经关闭的
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// 获取ws的写入器并发送信息，这个写入器需要调用它的close方法表示已经写入结束了
			// w, err := c.conn.NextWriter(websocket.BinaryMessage)
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				continue
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			// 逐个读取管道的消息数据
			n := len(c.send)
			// fmt.Println("Size of channel", len(c.send), cap(c.send))
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			// 调用close方法发送消息，如果报错则处理
			if err := w.Close(); err != nil {
				fmt.Println(err)
			}

		// 接收到定时器的消息，向前端的ws发送信息，检查一下连接是否正常
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 对于一名新的前端用户，建立一个websocket和Client与之通信
// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 建立websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		Hub_log.Println(err)
		return
	}

	// 新建Client并向hub注册
	client := &Client{
		hub:   hub,
		conn:  conn,
		send:  make(chan []byte, 256),
		index: hub.count,
		name:  clientNames[hub.count%len(clientNames)],
	}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
