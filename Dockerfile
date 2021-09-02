FROM golang:1.16.6-alpine
LABEL maintainer="Kyle Jackson <kyle@thepublicclouds.com>"

ENV DEV=true

WORKDIR $GOPATH/src/github.com/wizedkyle/sumocli
COPY . .
RUN chmod +x ./scripts/build.sh
RUN ./scripts/build.sh

ENTRYPOINT ["/go/bin/sumocli"]
