package controller

type AppController struct {
	HealthCheckController HealthCheckController
	CreateRoomController  CreateRoomController
	JoinRoomController    JoinRoomController
	WsRoomController      WsRoomController
}
