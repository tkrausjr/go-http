#!/bin/bash

log=/var/log/conference-app.log

export PATH=/usr/local/bin:/usr/local/sbin:/bin:/sbin:/usr/bin:/usr/sbin

exec &>>$log

( exec /go/src/go-http/bin/conference-app & >>$log ) &
exit 0
