FROM scratch

ADD configuration.txt    /configuration.txt
ADD weather-wear         /weather-wear

EXPOSE 8080

ENTRYPOINT ["/weather-wear"]
