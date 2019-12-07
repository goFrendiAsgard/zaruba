package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	filename := "test.json"
	data := map[string]int{}
	log.Println("BEGIN")
	file, err := os.Open(filename)
	if err != nil {
		defer file.Close()
	}
	// lock file
	syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	b, err := ioutil.ReadFile(filename)
	// read the data and modify
	log.Println("READ", string(b))
	if err = json.Unmarshal(b, &data); err != nil {
		log.Println("ERROR", err)
	}
	data["count"] = data["count"] + 1
	b, err = json.Marshal(data)
	if err != nil {
		log.Println("ERROR", err)
	}
	// write the data
	log.Println("WRITE", string(b))
	if err = ioutil.WriteFile(filename, b, 0644); err != nil {
		log.Println("ERROR", err)
	}
	// pura-pura kerja
	time.Sleep(time.Second)
	// unloack file
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}
