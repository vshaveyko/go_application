FROM golang:1.6

ADD . /go/src/server
RUN go get github.com/go-martini/martini
RUN go get github.com/martini-contrib/cors
RUN go get gopkg.in/pg.v4
RUN go install server
ENTRYPOINT /go/bin/server

EXPOSE 8080
