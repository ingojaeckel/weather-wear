#!/bin/sh

PREFIX=https://weather-wea.appspot.com

ab -n 1000 -c 2 ${PREFIX}/rest/forecast?cityId=5391997

echo 
