package websocket

import (
	"fmt"
	"inn/internal/gateway/model"
	"inn/internal/gateway/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type GateWayHandler struct {
	gs *service.GateWayService
}

func NewGateWayHandler(gs *service.GateWayService) *GateWayHandler {
	return &GateWayHandler{gs: gs}
}

func (gh *GateWayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go func(conn *websocket.Conn) {
		defer func() {
			//关闭连接
			//c.conn.Close()
		}()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}

			msgJson := gjson.ParseBytes(message)
			typ := msgJson.Get("type").Uint()
			data := msgJson.Get("data")

			fmt.Println("收到消息：", msgJson.String())

			switch typ {
			case 0: //心跳
				// uid := data.Get("uid").Uint()
				timeout := data.Get("timeout").Uint()
				conn.WriteJSON(map[string]interface{}{
					"type":    0,
					"timeout": timeout,
				})
			case 1: //上线消息
				uid := data.Get("uid").Uint()
				// 保存连接
				gh.gs.StoreConn(uid, &model.Conn{Uid: uid, Wsconn: conn})

				conn.WriteJSON(map[string]interface{}{
					"type":   1,
					"status": "success",
				})

			case 2: //查询消息
				// ownerUid := data.Get("ownerUid").Uint()
				// otherUid := data.Get("otherUid").Uint()
				// 通过rpc调用 messageService 查询消息
			case 3: //发送消息
				senderUid := data.Get("senderUid").Uint()
				recipientUid := data.Get("recipientUid").Uint()
				content := data.Get("content").String()
				msgType := data.Get("msgType").Int()
				gh.gs.SendMessage(senderUid, recipientUid, content, int32(msgType))
			case 5: //查询总未读
				//uid := data.Get("uid").Uint()
				// 通过rpc调用 messageService 查询总未读
			case 6: //处理ack
				//tid := data.Get("ownerUid").Uint()

			}
		}
	}(conn)
}