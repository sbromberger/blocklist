#!/bin/bash
set -e
BLOCKLIST=BLOCKLIST.txt
BLOCKLIST_PROG=./blocklist
BLOCKTMPSET=tmpblocks
# must be consistent with the yaml/blocklist
NETSET=block
IPSET=ipset
$BLOCKLIST_PROG > $BLOCKLIST
NGROUPS=$(tail -1 $BLOCKLIST | cut -d " " -f 2 | sed "s/$BLOCKTMPSET//")
NBLOCKS=$(wc -l $BLOCKLIST | cut -d " " -f 1)

NGROUPSEQ=$(seq $NGROUPS)

logger "Processing $NBLOCKS IP blocks into $NGROUPS groups"
# make sure the blocklist tables are present - set -e will force exit if error
for i in $NGROUPSEQ; do
  NETSETN=$NETSET$i
  $IPSET list $NETSETN > /dev/null
done

# create the temporary tables
for i in $NGROUPSEQ; do
  BLOCKTMPN=$BLOCKTMPSET$i
  $IPSET create $BLOCKTMPN nethash hashsize 32768 -exist
done

# restore the blocklist into the temp tables and swap the tables
$IPSET restore < $BLOCKLIST
for i in $NGROUPSEQ; do
  BLOCKTMPN=$BLOCKTMPSET$i
  NETSETN=$NETSET$i
  $IPSET swap $BLOCKTMPN $NETSETN -quiet
done

# delete the temporary tables
for i in $NGROUPSEQ; do
  BLOCKTMPN=$BLOCKTMPSET$i
  $IPSET flush $BLOCKTMPN -quiet
  $IPSET destroy $BLOCKTMPN -quiet
done
# log
logger $NGROUPS blocklists updated
