# ---------- 1단계: 빌드 ----------
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO를 끄면 외부 libc에 의존하지 않는 정적 바이너리가 나온다
RUN CGO_ENABLED=0 go build -o server .

# ---------- 2단계: 런타임 (최소) ----------
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 5000
ENTRYPOINT ["./server"]