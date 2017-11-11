#!/bin/bash
set -e
echo "[build.sh:building binary]"
cd $BUILDPATH && go build -o /api && rm -rf /tmp/*
echo "[build.sh:launching binary]"
/api