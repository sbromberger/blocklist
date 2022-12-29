#!/bin/bash
set -e
BLACKLIST="BLACKLIST.txt"
BLOCKLIST_PROG="./blocklist"
GEOTMPSET="geoip"
NETSET="GeoIP"
IPSET=echo
$BLOCKLIST_PROG > $BLACKLIST
NGROUPS=$(tail -1 $BLACKLIST | cut -d " " -f 2 | sed "s/geoip//")
echo $NGROUPS
for i in $(seq $NGROUPS); do
  GEOTMPN=$GEOTMPSET$i
  $IPSET create $GEOTMPN nethash hashsize 32768 -exist
done

$IPSET restore < $BLACKLIST
for i in $(seq $NGROUPS); do
  GEOTMPN=$GEOTMPSET$i
  NETSETN=$NETSET$i
  $IPSET swap $GEOTMPN $NETSETN -quiet && sudo $IPSET flush $GEOTMPN -quiet
done

# delete the temporary table
for i in $(seq $NGROUPS); do
  GEOTMPN=$GEOTMPSET$i
  $IPSET destroy $GEOTMPN -quiet
done
# log
echo logger $NGROUPS GeoIP blocklists updated.
