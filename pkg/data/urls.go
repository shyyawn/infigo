package data

import (
	"fmt"
	log "github.com/shyyawn/go-to/logging"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type Urls struct {
	Domains []struct {
		Alias  string   `yaml:"alias"`
		Domain string   `yaml:"domain"`
		Urls   []string `yaml:"urls"`
	} `yaml:"domains"`
}

func (u *Urls) LoadUrls(filePath string) *Urls {
	urlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error(fmt.Sprintf("yamlFile.Get err #%v ", err))
	}
	err = yaml.Unmarshal(urlFile, u)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unmarshal: %v", err))
	}
	return u
}

func (u *Urls) GetUrls(alias string) (domain string, urls []string) {
	for _, url := range u.Domains {
		if url.Alias == strings.ToLower(alias) {
			log.Warn(len(url.Urls))
			return url.Domain, url.Urls
		}
	}
	return "", nil
}
