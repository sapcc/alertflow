#!/bin/bash

export BILLING_ENDPOINT=https://billing.qa-de-1.cloud.sap:64000/
export BILLING_AUTH=

export SMTP_HOST=cronus.qa-de-1.cloud.sap:587
export SMTP_USER=
export SMTP_SECRET=

go run ../../main.go
