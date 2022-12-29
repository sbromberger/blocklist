#!/bin/bash
set -e
BLACKLIST=$(grep "^BLACKLIST" geoip.sh | cut -d "=" -f 2 | tr -d '"')
NGROUPS=$(tail -1 $BLACKLIST | cut -d " " -f 2 | sed "s/geoip//")
NETSET=$(grep "^NETSET" geoip.sh | cut -d "=" -f 2 | tr -d '"')
for i in $(seq $NGROUPS); do 
  sudo ipset test $NETSET$i $1
done
