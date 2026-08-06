# 1. 베이스 이미지는 `golang:1.22`를 사용한다
# 2. 작업 디렉터리는 `/app`으로 지정한다
# 3. `go.mod`를 먼저 복사하고, 그다음 소스를 복사한다
# 4. `go build`로 실행 바이너리를 만든다 (결과물 이름: `server`)
# 5. 컨테이너가 사용하는 포트(5000)를 명시한다
# 6. 컨테이너가 뜨면 자동으로 그 바이너리가 실행된다
# 7. 이미지 이름은 `counter:1.0`으로 빌드한다
# 8. 호스트 **8080** 포트로 접속하면 컨테이너 **5000** 포트로 연결된다

FROM golang:1.22
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o server

EXPOSE 5000
CMD ["./server"]