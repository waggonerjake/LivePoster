package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	steps := "Home,Home,Up,Select,Up,Select,Up,Up,Up,Select,Down,Select,Select" //Remote input to restart the Roku
	rokuAddress := fmt.Sprintf("http://%v:8060/keypress/", os.Args[1])          //Should be your IP of the Roku Device you want to restart
	file, err := os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	for _, element := range strings.Split(steps, ",") {
		action := fmt.Sprintf("%v%v", rokuAddress, element)
		req, err := http.NewRequest(http.MethodPost, action, nil)
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}
		resp, postErr := client.Do(req)
		if postErr != nil {
			log.Fatal(postErr)
		}
		log.Print(resp)

		time.Sleep(time.Second / 2)
	}
	time.Sleep(time.Second * 5)
	var resp *http.Response
	for i := 0; i < 5; i++ { //Try and verify up to 5 times
		fmt.Println(resp)
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%vHome", rokuAddress), nil) //Try and hit home to see if the reboot worked
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}
		resp, _ = client.Do(req)

		if resp != nil {
			break
		}
		time.Sleep(time.Second / 2)
	}
	defer file.Close()
	return
}
