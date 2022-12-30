#!/bin/bash
set -e
BLOCKLIST=$(grep "^BLOCKLIST=" runblocklist.sh | cut -d "=" -f 2 | tr -d '"')
BLOCKTMPSET=$(grep "^BLOCKTMPSET=" runblocklist.sh | cut -d "=" -f 2 | tr -d '"')
NGROUPS=$(tail -1 $BLOCKLIST | cut -d " " -f 2 | sed "s/$BLOCKTMPSET//")
NETSET=$(grep "^NETSET" runblocklist.sh | cut -d "=" -f 2 | tr -d '"')
for i in $(seq $NGROUPS); do 
  sudo ipset test $NETSET$i $1
done
