package socket

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

const ChatAction = "chat"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const DrawingAction = "drawing"

type Message struct {
	Action  string        `json:"action"`
	Target  MessageRoom   `json:"target"`
	Sender  MessageClient `json:"sender"`
	Payload interface{}   `json:"payload"`
}

type MessageClient struct {
	ClientName string `json:"client_name"`
	ClientID   string `json:"client_id"`
	AvatarUrl  string `json:"avatar_url"`
}

type MessageRoom struct {
	RoomID     string `json:"room_id"`
	MaxPlayers int    `json:"max_players"`
	Private    bool   `json:"private"`
}

type MessageJoinRoomPayload struct {
	Message string          `json:"message"`
	Clients []MessageClient `json:"clients"`
}

type MessageLeaveRoomPayload struct {
	Message string          `json:"message"`
	Clients []MessageClient `json:"clients"`
}

type MessageChatPayload struct {
	Message string `json:"message"`
}

type MessageDrawingPayload struct {
	OffsetX float64 `json:"offset_x"`
	OffsetY float64 `json:"offset_y"`
	State   string  `json:"state"`
}

func (message *Message) encode() []byte {
	json, err := json.Marshal(message)

	if err != nil {
		log.Error(err)
		return nil
	}

	return json
}
