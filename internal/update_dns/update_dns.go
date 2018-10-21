package update_dns

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"

	"gitlab.com/rafadc/ddns-update/internal/find_my_ip"
	"gitlab.com/rafadc/ddns-update/internal/namecheap"
)

type Domain struct {
	method    string
	domain    string
	subdomain string
	key       string
}

type Config struct {
	minutes_between_updates int
	domains                 map[string]Domain
}

var config = Config{}

func StartFromConfigFile() {
	readConfig()
	for {
		myIp, err := find_my_ip.MyIP()
		if err != nil {
			log.Printf("Couldn't get IP %s", err)
		} else {
			updateDomains(myIp)
		}
		time.Sleep(time.Duration(config.minutes_between_updates) * time.Minute)
	}
}

func updateDomains(myIp string) {
	for domainConfigName, domainDetails := range config.domains {
		log.Printf("Updating %s", domainConfigName)
		updateDns(domainDetails.method, domainDetails.subdomain, domainDetails.domain, myIp, domainDetails.key)
	}
}

func updateDns(method string, subdomain string, domain string, ip string, key string) {
	if method != "namecheap" {
		log.Fatal("Only namecheap supported")
		os.Exit(1)
	}
	namecheap.UpdateDns(subdomain, subdomain, ip, key)
}

func readConfig() {
	viper.SetConfigName("ddns-update")
	viper.AddConfigPath("/etc/ddns-update/")
	viper.AddConfigPath("$HOME/.ddns-update/")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault("minutes_between_updates", 20)
	viper.SetDefault("domains", map[string]Domain{})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config file: %s \n", err))
	}

	config = Config{
		minutes_between_updates: viper.Get("minutes_between_updates").(int),
		domains:                 convertMapOfStringsToMapOfDomains(viper.GetStringMap("domains")),
	}
}

func convertMapOfStringsToMapOfDomains(in map[string]interface{}) map[string]Domain {
	out := map[string]Domain{}
	for k, v := range in {
		out[k] = Domain{
			method:    v.(map[string]interface{})["method"].(string),
			domain:    v.(map[string]interface{})["domain"].(string),
			subdomain: v.(map[string]interface{})["subdomain"].(string),
			key:       v.(map[string]interface{})["key"].(string),
		}
	}
	return out
}
