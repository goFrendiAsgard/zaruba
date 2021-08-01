package advertisement

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"time"

	"gopkg.in/yaml.v3"
)

type Adv struct {
	Start   string `yaml:"start,omitempty"`
	End     string `yaml:"end,omitempty"`
	Pattern string `yaml:"pattern,omitempty"`
	Message string `yaml:"message,omitempty"`
}

type Advs map[string]Adv

func NewAdvs(fileName string) (advertisements *Advs, err error) {
	advertisements = &Advs{}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(b, advertisements); err != nil {
		return nil, fmt.Errorf("error parsing YAML '%s': %s", fileName, err)
	}
	return advertisements, err
}

func (advertisements *Advs) Get() string {
	format := "01-02-2006 15:04:05"
	candidates := []Adv{}
	currentTime := time.Now()
	currentPattern := currentTime.Format(format)
	for _, adv := range *advertisements {
		if adv.Start != "" {
			start, err := time.Parse(format, adv.Start)
			if err != nil || currentTime.Before(start) {
				continue
			}
		}
		if adv.End != "" {
			end, err := time.Parse(format, adv.End)
			if err != nil || currentTime.After(end) {
				continue
			}
		}
		if adv.Pattern != "" {
			match, err := regexp.Match(adv.Pattern, []byte(currentPattern))
			if err != nil || !match {
				continue
			}
		}
		candidates = append(candidates, adv)
	}
	if len(candidates) == 0 {
		return "Have a nice day!!!"
	}
	index := rand.Intn(len(candidates))
	candidate := candidates[index]
	return candidate.Message
}
