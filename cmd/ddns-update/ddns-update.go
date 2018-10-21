package main

import (
	"log"

	"github.com/rafadc/ddns-update/internal/update_dns"
)

func main() {
	log.Print("Starting...")
	update_dns.StartFromConfigFile()
}
