FROM golang:1.22
WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o server

RUN mkdir -p /data && \
    useradd -u 1000 -m appuser && \
    chown -R 1000:1000 /data
USER 1000

EXPOSE 5000
CMD ["./server"]