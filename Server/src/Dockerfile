FROM golang:1.16.5

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN ls

RUN pwd

RUN go mod download

RUN go build -o ./server .

EXPOSE 7014

CMD ["./server"]