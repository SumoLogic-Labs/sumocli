FROM golang:1.17.1-alpine
ARG buildnumber=DEV
ARG version=DEV
LABEL maintainer="Kyle Jackson <kyle@thepublicclouds.com>"

WORKDIR $GOPATH/src/github.com/SumoLogic-Incubator/sumocli
COPY . .
RUN chmod +x ./scripts/docker-build.sh
RUN ./scripts/docker-build.sh -b $buildnumber -v $version -d

ENTRYPOINT ["/go/bin/sumocli"]
