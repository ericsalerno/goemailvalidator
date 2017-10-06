package goemailvalidator

import (
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
		return 0
	}

	lines := strings.Split(string(dat), "\n")

	total := 0
	for i := 0; i < len(lines); i++ {
		host := lines[i]

		if host == "" {
			continue
		}

		if host[0:1] == "#" {
			continue
		}

		c.HostList[host] = 0

		total++
	}

	return total
}
