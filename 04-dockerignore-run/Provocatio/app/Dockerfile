FROM golang:1.22
WORKDIR /app
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

EXPOSE 5000
CMD [ "./server" ]