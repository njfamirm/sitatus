package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type site struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type siteList struct {
	List []site `yaml:"site-list"`
}

func GetSiteList() []site {
	data, err := ioutil.ReadFile("config/site-list.yml")
	if err != nil {
		fmt.Println(fmt.Errorf("(getSiteList): reading yaml file failed: %s", err))
	}

	var siteList siteList
	err = yaml.Unmarshal(data, &siteList)
	if err != nil {
		fmt.Println(fmt.Errorf("(getSiteList): unmarshalling YAML failed: %s", err))
	}

	return siteList.List
}
