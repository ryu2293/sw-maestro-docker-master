FROM golang:1.22
WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o server

RUN mkdir -p /data

EXPOSE 5000
CMD ["./server"]