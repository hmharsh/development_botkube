package main

import (
	"fmt"
	"github.com/hmharsh/botkube/pkg/config"
	log "github.com/hmharsh/botkube/pkg/log"
	"strings"
	//	"gopkg.in/yaml.v2"
)

func main() {

	conf, err := config.New()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error in loading configuration. Error:%s", err.Error()))
	}
	args := strings.Fields("get Pods")
	exists := authorizeCommandByProfile(conf.Communications.Slack.Accessbindings[0].ProfileValue, args)

	fmt.Printf("%+v\n", exists)
	//	// convert conf to proper string
	//	b, err := yaml.Marshal(conf)
	//	if err != nil {
	//		log.Fatal("unable to marshal")
	//
	//	}
	//	configYaml := string(b)
	//	fmt.Printf(configYaml)
	//	//fmt.Printf("%+v\n", conf.Communications.Mattermost)
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if strings.ToUpper(item) == strings.ToUpper(val) {
			return i, true
		}
	}
	return -1, false
}

func authorizeCommandByProfile(Profile config.Profile, args []string) bool {
	authorized_command := false

	if len(args) >= 1 {
		allowed_operations := Profile.Kubectl.Commands.Verbs
		if _, authorized_command = Find(allowed_operations, args[0]); !authorized_command {
			return false
		}
	}
	if len(args) >= 2 {
		allowed_resources := Profile.Kubectl.Commands.Resources
		if _, authorized_command = Find(allowed_resources, args[1]); !authorized_command {
			return false
		}
	}
	return authorized_command
}
