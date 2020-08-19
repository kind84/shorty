ARG GOVERSION=1.15.0
FROM golang:${GOVERSION}-alpine AS dep

ENV GOPROXY=https://proxy.golang.org

WORKDIR /shorty           

COPY go.mod .              
COPY go.sum .              

RUN go mod download 
