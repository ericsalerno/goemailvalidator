package main

import "testing"

func TestLoadBlacklist(t *testing.T) {
	c := Configuration{}

	c.LoadBlacklist("./test_data/test_blacklist.txt")

	if len(c.HostList) != 3 {
		t.Fatal("No items loaded in blacklist!")
	}

	if c.HostList["item1"] != 1 {
		t.Fatal("First item is incorrect!")
	}
}
