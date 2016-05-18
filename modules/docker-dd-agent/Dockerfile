FROM datadog/docker-dd-agent:11.0.574

ADD start.sh        /start.sh
ADD docker-dd-agent /background-service
EXPOSE 8080

CMD ["/start.sh"]

ENV PROXY_PORT 8080
ENV LOG_LEVEL DEBUG
ENV API_KEY fb75037c7f88c377d412c4130c650df9
