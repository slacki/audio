FROM golang:1.14-buster

RUN apt-get update -y && apt-get upgrade -y && apt-get install -y file

RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u github.com/githubnemo/CompileDaemon

WORKDIR /api

COPY . .

EXPOSE 8081
EXPOSE 40000

# docker-compose should run the actual build / CompileDaemon / delve 
