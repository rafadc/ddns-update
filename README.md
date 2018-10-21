# DDNS updater

A small tool to update your dynamic DNS. Currrently only Namecheap is supported.

## Config file

ddns-update will look for a config file called ddns-update.yaml in /etc/ddns-update/ in $HOME/.ddns-update or the current folder.

A sample config file would be

``` yaml
minutes_between_updates: 20

domains:
  mydomain.com:
    method: "namecheap"
    domain: "mydomain.com"
    subdomain: "rss"
    key: "<your ddns key goes here>"

  other_thing:
    method: "namecheap"
    domain: "mydomain.com"
    subdomain: "other_thing"
    key: "<your ddns key goes here>"

```

### Global configuration

 - *minutes_between_updates*: Number of minutes that pass between IP change checks.

### Domain list

There a list of domains is specified. The key is a descriptive name for the domain. The rest of the parameter are:

``` yaml
  mydomain.com:
    method: "namecheap"
    domain: "mydomain.com"
    subdomain: "rss"
    key: "<your ddns key goes here>"
```

 - *method*: The method to update the dns. At the moment only namecheap domains are supported.
 - *domain*: The top level domain name registered in Namecheap.
 - *subdomain*: The subdomain name to set up. It does not have to include the top level domain name. For example if you want to update rss.mydomain.com you only need to put rss here.
 - *key*: Your private ddns key

## Development

Following the guide at https://github.com/golang-standards/project-layout
