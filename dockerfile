FROM golang:1.16-buster AS builder
ARG DB_SERVER_NAME
ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD

ADD . /go/src/go-rest
WORKDIR /go/src/go-rest
#RUN go mod init go-rest
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
#RUN go mod tidy
RUN go get -u github.com/spf13/viper
# ENV DB_SERVER_NAME=db_server
# ENV DB_NAME=db_name
# ENV DB_USER=db_user
# ENV DB_PASSWORD=db_password
RUN go build -o /go/bin/go-rest
EXPOSE 8090
ENTRYPOINT ["/go/bin/go-rest"]