package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/netip"
	"sort"
	"strings"

	"go4.org/netipx"
)

type Config struct {
	MaxGroupSize int            `yaml:"maxgroupsize"`
	MaxGroups    int            `yaml:"maxgroups"`
	Prefix       string         `yaml:"prefix"`
	TempPrefix   string         `yaml:"tempprefix"`
	Ccs          []string       `yaml:"ccs"`
	OtherSources []string       `yaml:"othersources"`
	OtherBlocks  []netip.Prefix `yaml:"otherblocks"`
	WhiteList    []netip.Prefix `yaml:"whitelist"`
}

func getCountryURLs(countryCode string) []string {
	return []string{
		fmt.Sprintf("https://raw.githubusercontent.com/herrbischoff/country-ip-blocks/master/ipv4/%s.cidr", countryCode),
		fmt.Sprintf("https://www.ipdeny.com/ipblocks/data/aggregated/%s-aggregated.zone", countryCode),
	}
}

func clean(line string) (netip.Prefix, bool) {
	// strip leading and trailing spaces
	line = strings.TrimSpace(line)
	// ignore blank lines
	if len(line) == 0 {
		return netip.Prefix{}, false
	}
	// ignore if starts with #
	if strings.HasPrefix(line, "#") {
		return netip.Prefix{}, false
	}
	// ignore ipv6 addresses
	if strings.Contains(line, ":") {
		return netip.Prefix{}, false
	}
	if !strings.Contains(line, "/") {
		line = line + "/32"
	}
	// try parsing this as a CIDR address
	ipnet, err := netip.ParsePrefix(line)
	if err != nil {
		return netip.Prefix{}, false
	}
	return ipnet, true
}

func getAndClean(url string, ch chan netip.Prefix) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if cleaned, ok := clean(line); ok {
			ch <- cleaned
		}
	}
}

func (c *Config) getAllURLs() []netip.Prefix {
	var allURLs []string

	for _, cc := range c.Ccs {
		allURLs = append(allURLs, getCountryURLs(cc)...)
	}
	allURLs = append(allURLs, c.OtherSources...)

	ch := make(chan netip.Prefix, 1_000)
	remainingWorkers := len(allURLs)
	doneWorkers := make(chan bool, remainingWorkers)
	for _, url := range allURLs {
		go func(url string) {
			getAndClean(url, ch)
			doneWorkers <- true
		}(url)
	}

	var b netipx.IPSetBuilder
	for remainingWorkers > 0 {
		select {
		case prefix := <-ch:
			b.AddPrefix(prefix)
		case <-doneWorkers:
			remainingWorkers--
		}
	}
	close(doneWorkers)
	close(ch)

	// make sure the channel is empty
	for cidr := range ch {
		b.AddPrefix(cidr)
	}

	for _, cidr := range c.OtherBlocks {
		b.AddPrefix(cidr)
	}
	for _, cidr := range c.WhiteList {
		b.RemovePrefix(cidr)
	}

	prefixes, err := b.IPSet()
	if err != nil {
		return []netip.Prefix{}
	}
	pref := prefixes.Prefixes()
	sort.Slice(pref, func(i, j int) bool {
		return pref[i].Bits() < pref[j].Bits()
	})
	return pref
}
