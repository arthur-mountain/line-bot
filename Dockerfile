FROM golang:1.18

WORKDIR /usr/line-chatbot

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/bin/app 

CMD ["/usr/bin/app"]