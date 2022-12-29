# Blocklist - an easy way to add blocklists to your Edgerouter.

This package (a binary, a configuration file, and a shell script) can be used to download
blocklists based on country code and other arbitrary IP-based lists for use as filters on
Ubiquiti Edgerouters (and possibly other devices). Custom blocklist entries are supported,
as are custom whitelist entries. CIDR ranges are aggregated during processing, reducing
the complexity of the filter.

### Requirements:
- Go 1.19 or later (`go version` to determine version)

### Build and Installation (Edgerouter 4):
1. Clone the repo
2. Build on a system with Go installed for the edgerouter 4 architecture: `GOOS=linux GOARCH=mips go build`
3. Create a `geoip.yaml` configuration (use the example yaml for inspiration)
4. Edit the variables in geoip.sh to suit your setup (the defaults are sane but you might prefer different names)
5. Copy `blocklist`, `geoip.yaml`, and `geoip.sh` to your router (recommended: place in `/config/geoip`)
6. Test by running `geoip.sh` as root on the router
7. Create a cronjob to run periodically (optional)

