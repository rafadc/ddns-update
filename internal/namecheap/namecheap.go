package namecheap

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// UpdateDns changes the DNS entry for namecheap
func UpdateDns(subdomain string, domain string, ip string, key string) {
	url := "https://dynamicdns.park-your-domain.com/update?host=" + subdomain + "&domain=" + domain + "&password=" + key + "&ip=" + ip
	log.Printf("Updating %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("The server replied %s", body)
}
