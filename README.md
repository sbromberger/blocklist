# Blocklist - an easy way to add blocklists to your Edgerouter.

This package (a binary, a configuration file, and a shell script) can be used to download
blocklists based on country code and other arbitrary IP-based lists for use as filters on
Ubiquiti Edgerouters (and possibly other devices). Custom blocklist entries are supported,
as are custom whitelist entries. CIDR ranges are aggregated during processing, reducing
the complexity of the filter.

### Requirements:
- Go 1.19 or later

### Build and Installation (Edgerouter 4):
1. Clone the repo
2. Build for the edgerouter 4 architecture: `GOOS=linux GOARCH=mips go build`
3. Create a `geoip.yaml` configuration (use the example yaml for inspiration)
4. Copy `blocklist`, `geoip.yaml`, and `geoip.sh` to your router (recommend: place in `/config/geoip`)
5. Test by running `geoip.sh` as root on the router
6. Create a cronjob to run periodically.

