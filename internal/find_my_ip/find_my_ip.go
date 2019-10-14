package find_my_ip

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// MyIP retrieves your current IP address
func MyIP() (string, error) {
	ipFindSource := "http://myip.dnsomatic.com"
	resp, err := http.Get(ipFindSource)
	if err != nil {
		return "", errors.New("Can't get own ip from " + ipFindSource)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil || body != "" {
		return "", errors.New("Can't get own ip from " + ipFindSource)
	}

	log.Printf("Your IP is %s", string(body))

	return string(body), nil
}
