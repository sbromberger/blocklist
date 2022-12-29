package main

import (
	"fmt"
	"log"
	"net/netip"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var configFile string
	if len(os.Args) < 2 {
		configFile = "geoip.yaml"
	} else {
		configFile = os.Args[1]
	}

	f, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	config := Config{Ccs: []string{}, OtherSources: []string{}, OtherBlocks: []netip.Prefix{}, WhiteList: []netip.Prefix{}}
	if err := yaml.Unmarshal(f, &config); err != nil {
		log.Fatal(err)
	}
	ipnets := config.getAllURLs()
	if len(ipnets) > config.MaxGroupSize*config.MaxGroups {
		log.Fatalf("%d exceeds maximum size; increase maxGroups or maxGroupSize", len(ipnets))
	}
	// fmt.Printf("len(ipnets) = %d\n", len(ipnets))
	for i, ipnet := range ipnets {
		set := (i / config.MaxGroupSize) + 1
		var prefix string
		if config.Prefix == "" {
			prefix = fmt.Sprintf("add %s%d ", config.TempPrefix, set)
		} else {
			prefix = config.Prefix
		}
		fmt.Printf("%s%s\n", prefix, ipnet)
	}
}
