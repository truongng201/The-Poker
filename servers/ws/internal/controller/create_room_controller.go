package controller

import (
  utils "ws-service/pkg/utils"
  socket "ws-service/pkg/socket"
    
  "github.com/google/uuid"

  "github.com/labstack/echo/v4"
)

type CreateRoomController struct{}

type (
	CreateRoomReqBody struct {
		MaxPlayers int  `json:"max_players" validate:"required,min=2,max=20"`
		Private    bool `json:"private"`
	}
	CreateRoomResData struct {
		RoomID     string `json:"room_id"`
		MaxPlayers int    `json:"max_players"`
		Private    bool   `json:"private"`
	}
)

func (controller *CreateRoomController) Execute(c echo.Context, wsServer *socket.WsServer ) error {
  var reqBody CreateRoomReqBody
  if err := c.Bind(&reqBody); err != nil {
    return err
  }

  if err := c.Validate(&reqBody); err != nil {
    return err
  }

  roomID := uuid.New().String()
  wsServer.CreateRoom(roomID, reqBody.MaxPlayers, reqBody.Private)
  
	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Create Room Success",
	  Payload: &CreateRoomResData{
      RoomID:     roomID,
      MaxPlayers: reqBody.MaxPlayers,
      Private:    reqBody.Private,
    }
  })
} 
