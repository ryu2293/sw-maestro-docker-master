FROM golang:1.22
WORKDIR /api

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o server

EXPOSE 6000
CMD ["./server"]