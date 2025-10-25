package ws

import (
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/writer"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleDriverWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed : %v ", err)
		return
	}
	defer conn.Close()

	userID := c.Query("userID")
	if userID == "" {
		log.Println("No userID provided")
		return
	}

	packageSlug := c.Query("packageSlug")
	if packageSlug == "" {

		log.Println("No packageSlug provided")
		return
	}

	type Driver struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		ProfilePicture string `json:"profilePicture"`
		CarPlate       string `json:"carPlate"`
		PackageSlug    string `json:"packageSlug"`
	}

	msg := contracts.WSMessage{
		Type: "driver.cmd.location",
		Data: Driver{
			Id:             userID,
			Name:           "shivam",
			ProfilePicture: util.GetRandomAvatar(1),
			CarPlate:       "MH13VM",
			PackageSlug:    "Sedan",
		},
	}
	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("error sending a message from driver end %v", err)
		return
	}

	if err := conn.WriteMessage(1, []byte("Message yeh le extra right kia mene hehehe")); err != nil {
		log.Printf("error sending a message from driver end %v", err)
		return
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading messsage: %v ", message)
			writer.WriteError(c, http.StatusBadRequest, err.Error())
			break
		}

		log.Printf("recieved message: %s", message)

	}

}

func HandleRiderWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed : %v ", err)

		return
	}
	defer conn.Close()

	userID := c.Query("userID")
	if userID == "" {
		log.Println("No userID provided")
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading messsage: %v ", message)

			break
		}

		log.Printf("recieved message: %s", message)

	}
}
