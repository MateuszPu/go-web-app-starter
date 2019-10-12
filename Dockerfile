FROM golang:1.13-alpine
WORKDIR /webapp
RUN apk update
COPY . .

RUN go build -o bin/app

CMD ["/webapp/bin/app"]

EXPOSE 8080