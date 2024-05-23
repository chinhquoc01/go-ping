package main

import (
	"fmt"

	"net/http"

	"github.com/robfig/cron"
)

func main() {
	fmt.Println("start")
	c := cron.New()
	c.AddFunc("@every 30m", func() {
		_, err := http.Get("https://save-more-5pgp.onrender.com/")
		if err != nil {
			fmt.Println("error: ", err)
		}
	})

	c.Start()

	var forever chan struct{}
	<-forever
}
