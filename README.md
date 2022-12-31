# Blocklist - an easy way to add blocklists to your EdgeRouter.

This package (a binary, a configuration file, and a shell script) can be used to download
blocklists based on country code and other arbitrary IP-based lists for use as filters on
Ubiquiti EdgeRouters (and possibly other devices). Custom blocklist entries are supported,
as are custom whitelist entries. CIDR ranges are aggregated during processing, reducing
the complexity of the filter.

### Important Note
This package modifies your router configuration. Do not use if you are not well-versed
in building or installing software on your target hardware.

**This package has only been tested on the EdgeRouter 4.**

### Requirements:
- Go 1.19 or later (`go version` to determine version)

### Build and Installation:
1. Clone the repo
2. Create a `blocklist.yaml` configuration (use the example yaml for inspiration). Note the `maxgroups` value
3. Examine the `runblocklist.sh` shell script and customize the `NETSET` variable if desired (optional)
4. With the value of the `NETSET` variable, create `maxgroups` network groups (e.g, `block1` through `block4`)
5. On the router, create firewall rules that block sources from these network groups
6. Build on a system with Go installed (for the EdgeRouter 4 architecture: `GOOS=linux GOARCH=mips go build`)
7. Copy `blocklist`, `blocklist.yaml`, `testip.sh`, and `runblocklist.sh` to your router (recommended: place in `/config/blocklist`)
8. Test blocklist downloads by running `blocklist | wc -l` as root on the router and seeing how many entries would be created.
8. Test by running `runblocklist.sh` as root on the router
9. Create a cronjob to run periodically (optional)

### Determining whether an IP address is blocked
Use the `testip.sh` script to determine whether an IP address is in a blocklist:
`testip.sh <ip address>`

