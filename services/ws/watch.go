package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"meetingBooking/config"
)

func (c *ClientManager) Watch() {
	for {
		fmt.Println("--监听websocket通信--")
		select {
		case conn := <-Manager.OnLineUser:
			fmt.Printf("有新连接进入: %v\n", conn.ID)
			Manager.Clients[conn.ID] = conn //把连接放到用户管理上
			replyMsg := ReplyMessage{
				Code:    200,
				Content: "用户已连接到服务器",
			}

			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.OffLineUser:
			fmt.Printf("连接中断%s\n", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replayMsg := &ReplyMessage{
					Code:    200,
					Content: "用户与服务器中断",
				}

				msg, _ := json.Marshal(replayMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)                 //关闭通道
				delete(Manager.Clients, conn.ID) //删除用户管理中的指定ID
			}
		case broadcast := <-Manager.Broadcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendID //2<-1
			flag := false                     //默认对方是离线状态
			for id, conn := range Manager.Clients {
				if id != sendId { //
					continue
				}
				select {
				case conn.Send <- message: //对方是在线状态
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID //1->2
			fmt.Println(flag, "flag")
			if flag {
				replayMsg := &ReplyMessage{
					Code:    200,
					Content: "对方在线应答",
				}

				msg, _ := json.Marshal(replayMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err := insertMessage(config.MongoDbName, id, string(message), 1, int64(expirationTime)) //1表示消息已读
				if err != nil {
					fmt.Println("Error inserting message", err)
				}
			} else {
				replayMsg := &ReplyMessage{
					Code:    200,
					Content: "对方离线未应答",
				}

				msg, _ := json.Marshal(replayMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err := insertMessage(config.MongoDbName, id, string(message), 0, int64(expirationTime)) //0表示消息未读
				if err != nil {
					fmt.Println("Error inserting message", err)
				}
			}

		}
	}
}
