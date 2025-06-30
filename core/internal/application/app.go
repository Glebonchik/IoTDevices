package application

import (
	"IoTDevicesCore/internal/storage"
	"IoTDevicesCore/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Store  storage.Storage
	Config *config.Config
}

func (app *Application) RunApp(r *gin.Engine) error {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"server": "running"})
	})

	if err := r.Run(app.Config.Server.Port); err != nil {
		log.Fatal("Cannot run server:", err)
		return err
	}

	log.Print("Server running")

	return nil

}
