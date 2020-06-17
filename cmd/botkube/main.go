package main

import (
	"fmt"
	"github.com/hmharsh/botkube/pkg/config"
	log "github.com/hmharsh/botkube/pkg/log"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in loading configuration. Error:%s", err.Error()))
	}
	fmt.Printf("%+v\n", conf.Communications.Slack)
}
