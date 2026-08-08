FROM golang:1.22
WORKDIR /web

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o client

EXPOSE 5000
CMD ["./client"]