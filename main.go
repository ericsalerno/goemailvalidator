package main

func main() {
	config := Configuration{
		Port: 8081,
	}
	config.LoadBlacklist("blacklist.conf")

	service := Service{
		Config: &config,
	}

	service.Listen()
}
