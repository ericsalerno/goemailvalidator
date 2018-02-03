package main

import (
	"strings"
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)


func main() {
	viper.SetEnvPrefix("gemv")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// using standard library "flag" package
	flag.Int("port", 8081, "Port to listen on")
	flag.String("blacklist", "blacklist.conf", "Location of blacklist file")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.AutomaticEnv()
	portNumber := viper.GetInt("port") // retrieve value from viper
	blacklistFile := viper.GetString("blacklist") // retrieve value from viper


	config := Configuration{
		Port: portNumber,
	}
	config.LoadBlacklist(blacklistFile)

	service := Service{
		Config: &config,
	}

	service.Listen()
}
