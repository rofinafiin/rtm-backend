package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	rtpkg "github.com/rofinafiin/rtm-package"
	"github.com/whatsauth/whatsauth"
	"net/http"
	"rtm/config"
)

var usercol = "data_user"

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetHome(c *fiber.Ctx) error {
	//getip := musik.GetIPaddress()
	getip := "Hello guys"
	return c.JSON(getip)
}

func Getdatauser(c *fiber.Ctx) error {
	id := "cc2"
	getstats := rtpkg.GetDatauser(id, config.MongoConn, usercol)
	fmt.Println("Data user : ", getstats)
	return c.JSON(getstats)
}

func InsertData(c *fiber.Ctx) error {
	database := config.MongoConn
	var jumlah rtpkg.User
	if err := c.BodyParser(&jumlah); err != nil {
		return err
	}
	Inserted := rtpkg.InsertDataUser(database,
		jumlah.Iduser,
		jumlah.Nama,
		jumlah.Email,
		jumlah.Handphone,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": Inserted,
	})
}

func GetDataUserbyPhone(c *fiber.Ctx) error {
	hp := c.Params("handphone")
	data := rtpkg.GetDataUserFromPhone(hp, config.MongoConn, "data_user")
	fmt.Println("data user phone : ", data)
	return c.JSON(data)
}

func DeleteDataUser(c *fiber.Ctx) error {
	hp := c.Params("handphone")
	data := rtpkg.DeleteData(hp, config.MongoConn, "data_user")
	return c.JSON(data)
}
