package service

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

type server struct{}

func (p *server) Start(s service.Service) error {
	return nil
}

func (p *server) Stop(s service.Service) error {
	return nil
}

func serverService() service.Service {

	var args []string

	if len(os.Args) > 5 {
		args = append(os.Args[3:4], os.Args[5:]...)
	}

	svcConfig := &service.Config{
		Name:        "tdp-cloud",
		DisplayName: "tdp cloud server",
		Description: "tdp cloud server",
		Arguments:   args,
	}

	s, err := service.New(&server{}, svcConfig)

	if err != nil {
		log.Fatal("init service error:", err)
	}

	return s

}

func serverInstall() {

	s := serverService()

	if x := s.Install(); x != nil {
		log.Print("install service error:", x.Error())
	} else {
		log.Print("install service done")
	}

}

func serverUninstall() {

	s := serverService()

	if x := s.Uninstall(); x != nil {
		log.Print("uninstall service error:", x.Error())
	} else {
		log.Print("uninstall service done")
	}

}