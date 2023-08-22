#!/bin/bash

CURL='/usr/bin/curl'
URL="http://127.0.0.1:8080/alerts"
# HEADER1="User-Agent: Alertmanager/0.23.0"
# HEADER2="Content-Type: application/json"
REQ="POST"
DATA="@alert.json"
#-H $HEADER1 -H $HEADER2
raw="$($CURL --request $REQ --data $DATA $URL)"
echo $raw
