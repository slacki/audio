FROM golang:1.14-buster

RUN apt update -y && apt upgrade -y && apt install file

RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u github.com/githubnemo/CompileDaemon

WORKDIR /api

COPY . .

EXPOSE 8081
EXPOSE 40000

# docker-compose should run the actual build / CompileDaemon / delve 
