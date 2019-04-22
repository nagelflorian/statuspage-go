FROM circleci/golang:1.12 AS base
WORKDIR /src

FROM base
USER root
RUN go get -u golang.org/x/lint/golint
COPY ./ ./
RUN golint ./...
RUN go test -mod vendor -v 2>&1 ./...

FROM base AS build
USER root
COPY ./ ./
RUN go build -mod vendor