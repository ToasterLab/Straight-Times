FROM golang:1.12.7-alpine3.10

ADD . /app
WORKDIR /app
RUN go build -o main .

CMD ["/app/main"]