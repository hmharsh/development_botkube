package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type conf struct {
	Hits             int64  `yaml:"hits"`
	Time             int64  `yaml:"time"`
	Myprofile        string `yaml:"myprofilee"`
	Selected_profile Profile
}
type AllProfiles struct {
	Profiles []Profile
}

func (all_profiles AllProfiles) getProfile(profile_name string) (Profile, error) {
	p := Profile{}
	for _, profile := range all_profiles.Profiles {
		if profile.Nam == profile_name {
			p = profile
			return p, nil
		}
	}
	return p, errors.New("selected profile not found in provided profiles")
}

type Profile struct {
	Nam         string `yaml:"nam"`
	Description string `yaml:"description"`
}

func (c *conf) getConf() *conf {
	p := &AllProfiles{}
	yamlFile, err := ioutil.ReadFile("/home/harshit/project/development/test/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	yamlFile, err = ioutil.ReadFile("/home/harshit/project/development/test/profile.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, p)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	c.Selected_profile, err = (p.getProfile(c.Myprofile))
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	var c conf
	c.getConf()

	fmt.Printf("\n%+v\n", c)
}
