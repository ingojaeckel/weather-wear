#!/bin/sh
/background-service &
/entrypoint.sh supervisord -n -c /etc/dd-agent/supervisor.conf
