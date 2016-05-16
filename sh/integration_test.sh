#!/bin/sh

PREFIX=https://weather-wea.appspot.com

curl -XGET ${PREFIX}/rest/forecast?cityId=5391997

echo 
