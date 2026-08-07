FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN RUN CGO_ENABLED=0 go build -o server .


FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 5000
ENTRYPOINT [ "./server" ]