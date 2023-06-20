FROM golang:alpine
WORKDIR /go/src
COPY web.go .
COPY go.mod .
RUN go build -o /go/web
expose 8888
CMD ["/go/web"]


