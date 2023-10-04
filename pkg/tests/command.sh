#!/bin/bash

CURL='/usr/bin/curl'
URL="http://127.0.0.1:8080/alerts"
REQ="POST"
DATA="@alert.json"

raw="$($CURL --request $REQ --data $DATA $URL)"
echo $raw
