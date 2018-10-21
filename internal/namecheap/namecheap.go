package namecheap

import (
	"log"
	"net/http"
	"os"
)

func UpdateDns(subdomain string, domain string, ip string, key string) {
	url := "https://dynamicdns.park-your-domain.com/update?host=" + subdomain + "&domain=" + domain + "&password=" + key + "&ip=" + ip
	log.Printf("Updating %s", url)
	_, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
