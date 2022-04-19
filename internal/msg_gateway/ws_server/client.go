package ws_server

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn
	// 用户ID
	userID string
	// 平台ID
	platformID int32
	// 上线时间戳（毫秒）
	onlineAt int64
	// Buffered channel of outbound messages.
	send chan []byte
	// 关闭通知
	close  chan []byte
	closed bool
	sync.Mutex
}

func newClient(hub *Hub, conn *websocket.Conn, userID string, platformID int32) *Client {
	return &Client{
		hub:        hub,
		conn:       conn,
		userID:     userID,
		platformID: platformID,
		onlineAt:   time.Now().UnixNano() / 1e6,
		send:       make(chan []byte, WsChanClientSendMessage),
		close:      make(chan []byte),
	}
}

func (c *Client) closeConn() {
	c.Lock()
	if c.closed {
		c.Unlock()
		return
	}
	c.closed = true
	c.Unlock()

	c.conn.Close()
	close(c.send)
	close(c.close)

	c.hub.unregister <- c
}

func (c *Client) read() {
	defer func() {
		c.closeConn()
	}()

	var (
		msgType   int
		bufMsg    []byte
		bufHeader []byte
		msgCode   int
		message   *Message
		err       error
	)

	c.conn.SetReadLimit(WsMaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(WsPongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(WsPongWait)); return nil })

	for {
		if msgType, bufMsg, err = c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//TODO:需要添加日志
				log.Printf("error: %v", err)
			}
			break
		}
		if msgType == websocket.PingMessage {
			continue
		}
		if len(bufMsg) == 0 {
			continue
		}
		bufHeader = bufMsg[0:WsHeaderLength]
		msgCode = bytes2Int(bufHeader)
		if msgCode == WsMsgCodePing {
			c.Send(WsMsgBufPong)
			continue
		}
		message = &Message{
			Client: c,
			Body:   bufMsg,
		}
		c.hub.read <- message
	}
}

func (c *Client) write() {
	pingTicker := time.NewTicker(WsPingPeriod)
	defer func() {
		pingTicker.Stop()
		c.closeConn()
	}()

	var (
		err     error
		message []byte
		ok      bool
	)

	for {
		select {
		case message, ok = <-c.send:
			if ok == false {
				// chan 关闭
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(WsWriteWait)); err != nil {
				c.conn.WriteMessage(websocket.CloseMessage, WsMsgBufClose)
				return
			}
			if err := c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
				return
			}
		case <-pingTicker.C:
			c.conn.SetWriteDeadline(time.Now().Add(WsWriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-c.close:
			return
		}
	}
}

func (c *Client) Send(message []byte) {
	if c.closed {
		return
	}
	c.send <- message
}

func (c *Client) Close() {
	if c.closed {
		return
	}
	c.close <- WsMsgBufClose
}