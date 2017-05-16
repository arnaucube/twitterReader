package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func festivalEn(text string) {
	path, err := exec.LookPath("festival")
	if err != nil {
		return
	}
	festival := exec.Command(path, "--tts")
	reader := strings.NewReader(text)
	festival.Stdin = reader
	err = festival.Run()
}

func festivalCa(text string) {
	out, err := exec.Command("bash", "readCa.sh", text).Output()
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(out))
}
