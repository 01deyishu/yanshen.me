#!/bin/bash
#get somp snmp info
snmpwalk -v 2c -c 7Niuread 100.64.0.254 1.3.6.1.2.1.2.2.1.2.$1 | awk '{print $NF}'