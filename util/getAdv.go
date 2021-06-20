package util

import (
	"github.com/state-alchemists/zaruba/advertisement"
)

func GetAdv(advertisementFile string) (message string, err error) {
	advs, err := advertisement.NewAdvs(advertisementFile)
	if err != nil {
		return "", err
	}
	return advs.Get(), nil
}
