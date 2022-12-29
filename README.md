# Blocklist - an easy way to add blocklists to your Edgerouter.

This package (a binary, a configuration file, and a shell script) can be used to download
blocklists based on country code and other arbitrary IP-based lists for use as filters on
Ubiquiti Edgerouters (and possibly other devices). Custom blocklist entries are supported,
as are custom whitelist entries. CIDR ranges are aggregated during processing, reducing
the complexity of the filter.

### Important Note
This package modifies your router configuration. Do not use if you are not well-versed
in building or installing software on your target hardware.

### Requirements:
- Go 1.19 or later (`go version` to determine version)

### Build and Installation (Edgerouter 4):
1. Clone the repo
2. Create a `geoip.yaml` configuration (use the example yaml for inspiration). Note the `maxgroups` value
3. Examine the `geoip.sh` shell script and customize the `NETSET` variable if desired
4. With the value of the `NETSET` variable, create `maxgroups` network groups (e.g, `GeoIP1` through `GeoIP4`)
5. On the router, create firewall rules that block sources from these network groups
6. Build on a system with Go installed for the edgerouter 4 architecture: `GOOS=linux GOARCH=mips go build`
7. Edit the variables in geoip.sh to suit your setup (the defaults are sane but you might prefer different names)
8. Copy `blocklist`, `geoip.yaml`, `testmulti.sh`, and `geoip.sh` to your router (recommended: place in `/config/geoip`)
9. Test by running `geoip.sh` as root on the router
10. Create a cronjob to run periodically (optional)

### Determining whether an IP address is blocked
Use the `testip.sh` script to determine whether an IP address is in a blocklist:
`testip.sh <ip address>`

