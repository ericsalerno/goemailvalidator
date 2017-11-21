package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Configuration for a service
type Configuration struct {
	Port     int
	HostList map[string]int
}

// LoadBlacklist load invalid hosts file
// Example https://raw.githubusercontent.com/martenson/disposable-email-domains/master/disposable_email_blacklist.conf
func (c *Configuration) LoadBlacklist(filename string) int {
	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Can not open file " + filename)
		return 0
	}

	lines := strings.Split(string(dat), "\n")
	c.HostList = make(map[string]int)

	total := 0
	for i := 0; i < len(lines); i++ {
		host := strings.TrimSpace(lines[i])
		host = strings.ToLower(host)

		if host == "" {
			continue
		}

		if host[0:1] == "#" {
			continue
		}

		c.HostList[host] = 1

		total++
	}
	fmt.Printf("Loaded %d blacklisted domains.\n", total)

	return total
}
