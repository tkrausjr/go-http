# Dockerfile used to download latest Source for Conference App
# and build a new Docker Image with default Entrypoint
# Listening on port TCP 8080

FROM golang:1.8 

MAINTAINER  TKraus 

RUN cd /go/src/ && \
	git clone https://github.com/tkrausjr/go-http.git

RUN cd /go/src/go-http/ && \
	go build -x -o ./bin/conference-app conference-app.go

ENTRYPOINT ["/go/src/go-http/conference-app-wrapper.sh"]
