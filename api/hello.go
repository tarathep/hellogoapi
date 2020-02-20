package api

import (
	"fmt"

	"github.com/tarathep/hellogoapi/event"

	"github.com/tarathep/hellogoapi/repository"

	"github.com/gin-gonic/gin"
	"github.com/tarathep/hellogoapi/models"
)

type HelloHandler struct {
	DB    repository.HelloLanguage
	Kafka event.Producer
}

func (h *HelloHandler) GetHello(c *gin.Context) {

	hellos, err := h.DB.AllHello()
	if err != nil {
		return
	}
	c.JSON(200, hellos)
}

func (h *HelloHandler) PostHello(c *gin.Context) {

	hello := models.Hello{}

	if err := c.ShouldBindJSON(&hello); err != nil {
		return

	}

	result, err := h.DB.InsertHello(hello)
	h.Kafka.Publish("test", "message1")

	//kafka.Producer{}.Publish("10.138.36.165:9092", "test", "message1")
	fmt.Println("Inserted a single document: ", result, err)

	c.String(200, "success")
}
