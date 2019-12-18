package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Service struct {
	Location string   `json:"location"`
	Response Response `json:"response"`
}

type Response struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	body    string            `json:"body"`
}

var services = []Service{
	{
		Location: "/webhook/postStage",
		Response: Response{
			Status: 200,
			body:   "OK",
		},
	},
	{
		Location: "/webhook/postUpgrade",
		Response: Response{
			Status: 200,
			body:   "OK",
		},
	},
	{
		Location: "/webhook/upgrading",
		Response: Response{
			Status: 200,
			body:   "OK",
		},
	},
}

var (
	logger         *log.Logger
	serverBindAddr string
)

func main() {
	logger = log.New(os.Stderr, "server", log.Flags())

	flag.StringVar(&serverBindAddr, "bind-addr", ":8080", "server bind addr")

	for _, service := range services {
		http.HandleFunc(service.Location, func(writer http.ResponseWriter, request *http.Request) {
			reqBody, err := ioutil.ReadAll(request.Body)
			if err != nil {
				logger.Println(err)
			}
			logger.Println(string(reqBody))
			writer.WriteHeader(service.Response.Status)
			_, err = writer.Write([]byte(service.Response.body))
			if err != nil {
				logger.Println(err)
			}
		})
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Panic(err)
	}
}
