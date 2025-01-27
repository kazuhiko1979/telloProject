package main

import (
	"log"
	"telloProject/app/controllers"
	"telloProject/config"
	"telloProject/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println(controllers.StartWebServer())

	// droneManager := models.NewDroneManager()
	// droneManager.TakeOff()
	// time.Sleep(10 * time.Second)
	// droneManager.Patrol()
	// time.Sleep(30 * time.Second)
	// droneManager.Patrol()
	// time.Sleep(10 * time.Second)
	// droneManager.Land()
}
