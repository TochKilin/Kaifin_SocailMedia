package config

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Client struct {
	Conn *websocket.Conn
	ID   string
}

// ======================
// WebSocket Manager
// ======================
type WebSocketManager struct {
	mu      sync.RWMutex
	clients map[string]*Client
}

func ConnectPostgres(ctx context.Context) (*sqlx.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ======================
// SINGLETON INSTANCE
// ======================
var manager = NewWebSocketManager()

func GetWebSocketManager() *WebSocketManager {
	return manager
}

// ======================
// CREATE MANAGER
// ======================
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(map[string]*Client),
	}
}

// ======================
// ADD CLIENT
// ======================
func (wm *WebSocketManager) AddClient(client *Client) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	wm.clients[client.ID] = client
	log.Printf("Client added: %s", client.ID)
}

// ======================
// REMOVE CLIENT
// ======================
func (wm *WebSocketManager) RemoveClient(clientID string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	delete(wm.clients, clientID)
	log.Printf("Client removed: %s", clientID)
}

// ======================
// DEBUG CLIENTS
// ======================
func (wm *WebSocketManager) PrintClients() {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	log.Printf("Clients: %+v", wm.clients)
}

// ======================
// BROADCAST TO ALL
// ======================
func (wm *WebSocketManager) Broadcast(data interface{}) {

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

	for _, client := range clients {

		err := client.Conn.WriteJSON(data)
		if err != nil {
			log.Printf("Broadcast error: %v", err)
			client.Conn.Close()
			wm.RemoveClient(client.ID)
		}
	}
}

// ======================
// SEND TO ONE CLIENT
// ======================
func (wm *WebSocketManager) Emit(clientID string, data interface{}) {

	wm.mu.RLock()
	client, ok := wm.clients[clientID]
	wm.mu.RUnlock()

	if !ok {
		log.Printf("Client %s not found", clientID)
		return
	}

	err := client.Conn.WriteJSON(data)
	if err != nil {
		log.Printf("Emit error: %v", err)
		client.Conn.Close()
		wm.RemoveClient(clientID)
	}
}

// ======================
// SEND TO USER
// ======================
func (wm *WebSocketManager) NotifyUser(userID string, data interface{}) {

	clientID := "user-" + userID

	wm.mu.RLock()
	client, ok := wm.clients[clientID]
	wm.mu.RUnlock()

	if !ok {
		log.Printf("User client not connected: %s", clientID)
		return
	}

	err := client.Conn.WriteJSON(data)
	if err != nil {
		log.Printf("Notify error: %v", err)
		client.Conn.Close()
		wm.RemoveClient(clientID)
	}
}
