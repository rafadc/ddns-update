package find_my_ip

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func MyIP() (string, error) {
	ip_find_source := "http://myip.dnsomatic.com"
	resp, err := http.Get(ip_find_source)
	if err != nil {
		return "", errors.New("Can't get own ip from " + ip_find_source)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil || body != "" {
		return "", errors.New("Can't get own ip from " + ip_find_source)
	}

	log.Printf("Your IP is %s", string(body))

	return string(body), nil
}
