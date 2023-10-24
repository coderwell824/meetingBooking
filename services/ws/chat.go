package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"meetingBooking/config"
	"meetingBooking/repository/cache"
	"net/http"
	"strconv"
	"time"
)

const expirationTime = time.Hour * 24 * 30 * 3 //过期时间为3个月

// SendMessage 发送消息的结构体
type SendMessage struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ReplyMessage 回复消息的结构体
type ReplyMessage struct {
	MsgFrom string `json:"MsgFrom"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

// ClientUser 客户端用户的结构体
type ClientUser struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

// Broadcast 广播结构体
type Broadcast struct {
	Client  *ClientUser
	Message []byte
	Type    int
}

// ClientManager 用户管理类
type ClientManager struct {
	Clients     map[string]*ClientUser
	Broadcast   chan *Broadcast
	Reply       chan *ClientUser
	OnLineUser  chan *ClientUser
	OffLineUser chan *ClientUser
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:     make(map[string]*ClientUser), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:   make(chan *Broadcast),
	OnLineUser:  make(chan *ClientUser),
	Reply:       make(chan *ClientUser),
	OffLineUser: make(chan *ClientUser),
}

func formatSendId(formId, toUid string) string {
	return formId + "->" + toUid
}

func ConnectionWsService(ctx *gin.Context, uId, toUid string) {

	//升级成ws协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { //CheckOrigin用来解决跨域问题
			return true
		}}).Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println("升级ws协议失败", err)
		return
	}

	//创建用户实例
	client := &ClientUser{
		ID:     formatSendId(uId, toUid),
		SendID: formatSendId(toUid, uId),
		Socket: conn,
		Send:   make(chan []byte),
	}
	//用户注册到用户管理上
	Manager.OnLineUser <- client

	go client.Read()
	go client.Write()
}

func (c *ClientUser) Read() {
	defer func() {
		Manager.OffLineUser <- c
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMessage)

		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			log.Println("Error reading send message", err)
			Manager.OffLineUser <- c
			_ = c.Socket.Close()
			break
		}

		if sendMsg.Type == 1 { //发送消息
			user1, _ := cache.RedisClient.Get(c.ID).Result()
			user2, _ := cache.RedisClient.Get(c.SendID).Result()

			if user1 > "3" && user2 == "" { //用户1给用户2发送了三条消息， 对方未回复
				replyMsg := ReplyMessage{
					Code:    403,
					Content: "对方未回复，等待对方回复后发送",
				}

				msg, _ := json.Marshal(replyMsg) //序列化
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				cache.RedisClient.Incr(c.ID)
				//防止websocket断开连接， 建立连接后三个月过期
				_, _ = cache.RedisClient.Expire(c.ID, expirationTime).Result()
			}

			//将消息进行广播出去
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content), //发送过来的消息
			}

		} else if sendMsg.Type == 2 { //获取历史消息
			timeT, err := strconv.Atoi(sendMsg.Content)
			if err != nil {
				timeT = 999999
			}
			historyMessage, _ := FindMany(config.MongoDbName, c.SendID, c.ID, int64(timeT), 10) //获取10条历史消息
			fmt.Println(historyMessage, "historyMessage")
			if len(historyMessage) > 10 {
				historyMessage = historyMessage[:10]
			} else if len(historyMessage) == 0 {
				replyMessage := ReplyMessage{
					Code:    200,
					Content: "暂无最新消息",
				}
				msg, _ := json.Marshal(replyMessage)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			}
			for _, msg := range historyMessage {
				replyMessage := ReplyMessage{
					MsgFrom: msg.From,
					Content: fmt.Sprintf("%v", msg.Content),
				}
				msg, _ := json.Marshal(replyMessage)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		} else if sendMsg.Type == 3 { //第一条消息
			firstMsg, err := FirstFindMsg(config.MongoDbName, c.SendID, c.ID)

			if err != nil {
				log.Println("Error finding message", err)
			}

			for _, msg := range firstMsg {
				replyMessage := ReplyMessage{
					MsgFrom: msg.From,
					Content: fmt.Sprintf("%v", msg.Content),
				}
				msg, _ := json.Marshal(replyMessage)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}

		}

	}
}

func (c *ClientUser) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			replyMsg := ReplyMessage{
				Code:    200,
				Content: string(message),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
