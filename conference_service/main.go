package main

import (

	"github.com/spf13/viper"
	"github.com/menachem554/Go-conference/conference_service/config"
	"github.com/menachem554/Go-conference/conference_service/server"
	service "github.com/menachem554/Go-conference/conference_service/service"
)

func main() {
	config.SetConfig()
	conference := service.CreateConferenceService()
	server := server.NewServer(&conference, viper.GetString("Port"))
	server.Serve()	
}