package main

import (
	"github.com/akerl/grafana-json/server"
)

func main() {
	s := server.JSONServer{}
	s.Start()
}
