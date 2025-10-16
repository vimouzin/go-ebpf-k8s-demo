FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o server .

EXPOSE 8081

CMD ["./server"]

