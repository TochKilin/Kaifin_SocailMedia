package websocket

import (
	// standard library
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	//interal package

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn *websocket.Conn
	ID   string
}

type WebSocketManager struct {
	mu      sync.RWMutex
	clients map[string]*Client
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(map[string]*Client),
	}
}

func (wm *WebSocketManager) PrintClients() {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	fmt.Println("clients:", wm.clients)
}

// =====================
// ADD CLIENT
// =====================
func (wm *WebSocketManager) AddClient(c *Client) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	wm.clients[c.ID] = c
	log.Println("client added:", c.ID)
}

// =====================
// REMOVE CLIENT
// =====================
func (wm *WebSocketManager) RemoveClient(id string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	delete(wm.clients, id)
	log.Println("client removed:", id)
}

// =====================
// BROADCAST
// =====================
func (wm *WebSocketManager) Broadcast(data any) {

	// copy clients safely
	wm.mu.RLock()
	clients := make([]*Client, 0, len(wm.clients))
	for _, c := range wm.clients {
		clients = append(clients, c)
	}
	wm.mu.RUnlock()

	// optional base64 decode
	if str, ok := data.(string); ok {
		if decoded, err := base64.StdEncoding.DecodeString(str); err == nil {
			data = string(decoded)
		}
	}

	// send outside lock
	for _, client := range clients {
		if err := client.Conn.WriteJSON(data); err != nil {
			log.Println("broadcast error:", err)
			client.Conn.Close()
			wm.RemoveClient(client.ID)
		}
	}
}

// =====================
// SEND TO ONE
// =====================
func (wm *WebSocketManager) Emit(id string, data any) {

	wm.mu.RLock()
	client, ok := wm.clients[id]
	wm.mu.RUnlock()

	if !ok {
		return
	}

	if err := client.Conn.WriteJSON(data); err != nil {
		log.Println("emit error:", err)
		client.Conn.Close()
		wm.RemoveClient(id)
	}
}

// =====================
// SEND TO USER
// =====================
func (wm *WebSocketManager) NotifyUser(userID string, data any) {

	id := "user-" + userID

	wm.mu.RLock()
	client, ok := wm.clients[id]
	wm.mu.RUnlock()

	if !ok {
		log.Println("user not connected:", id)
		return
	}

	if err := client.Conn.WriteJSON(data); err != nil {
		log.Println("notify error:", err)
		client.Conn.Close()
		wm.RemoveClient(id)
	}
}
