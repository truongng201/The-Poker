package controller

import (
	config "ws-service/config"
	socket "ws-service/pkg/socket"
	utils "ws-service/pkg/utils"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type WsRooomRequestBody struct {
	RoomId string `json:"room_id"`
}

type WsRoomController struct{}

func (controller *WsRoomController) Execute(c echo.Context, wsServer *socket.WsServer) error {
	var req WsRooomRequestBody
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}
	var upgrader = config.WsUpgrade()
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	defer conn.Close()

	log.Info("Client address: ", conn.RemoteAddr())

	client := socket.NewClient(
		conn, "", "", wsServer,
	)

	go client.WriteMessage()
	client.ReadMessage()

	wsServer.Register <- client

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Close connection",
		Payload: "",
	})
}
