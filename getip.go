package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func GetIp() string {
	ww, err := exec.Command("CMD", "/C", " ipconfig").Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	return fmt.Sprintf("%s", reg.FindAllString(string(ww), -1)[0])
}
