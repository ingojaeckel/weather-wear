# weather-wear

[![Build Status](https://travis-ci.org/ingojaeckel/weather-wear.svg?branch=master)](https://travis-ci.org/ingojaeckel/weather-wear)

A weather forecast app that tells you what is really important: What to wear.

# Starting the service locally
* `$ git clone git@github.com:ingojaeckel/weather-wear.git`
* `$ cd weather-wear`
* Create an account with [openweathermap.org](http://openweathermap.org/). Download your API key and insert it into the `configuration.txt` file
* `$ ./sh/run`

# Querying the service
* Find out the city ID by retrieving it via [http://openweathermap.org/current](http://openweathermap.org/current).
* Retrieve a recommendation on what to wear via `curl`:
* `$ curl https://weather-wea.appspot.com/rest/forecast?cityId=5391997` to run against the deployed service
* `$ curl http://localhost:8080/rest/forecast?cityId=5391997` to run against the locally running service
