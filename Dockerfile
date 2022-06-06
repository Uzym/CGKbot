FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get github.com/lib/pq
RUN go get github.com/Syfaro/telegram-bot-api
RUN go get gopkg.in/yaml.v3

RUN go build -o main .

CMD ["/app/main"]

