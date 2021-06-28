package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Generate(templateLocation, destination string, replacementMap map[string]string) {
	err := filepath.Walk(templateLocation,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size(), info.Mode())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
