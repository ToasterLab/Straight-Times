FROM golang:1.12.7-alpine3.10

ADD . /app
WORKDIR /app
RUN apk add --no-cache git mercurial \
    && go get -u github.com/labstack/echo/... \
    && apk del git mercurial
RUN go build -o main .

CMD ["/app/main"]
