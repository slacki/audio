FROM golang

WORKDIR /api

# delve doesn't work atm
# https://github.com/go-delve/delve/issues/1988
#RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get github.com/githubnemo/CompileDaemon

COPY . .

# # running this container as non-existent user will cause `go build` to fail if this is not defined
# # when go tries to write to default ~/.cache folder, the permissions are denied
# RUN mkdir /.gocache
# ENV GOCACHE /.gocache
# RUN chmod -R 777 /.gocache
# RUN chmod -R 777 /go

EXPOSE 8081
EXPOSE 40000

# docker-compose should run the actual build / CompileDaemon / delve 
