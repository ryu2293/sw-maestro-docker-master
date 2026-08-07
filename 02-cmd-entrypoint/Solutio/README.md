# 해설 — 문제 02

## 1. Dockerfile 정답

```dockerfile
FROM golang:1.22
WORKDIR /app
ENTRYPOINT ["./server"]
# 기본 인자 — docker run 뒤에 인자를 주면 이 줄이 덮어써진다
CMD ["-msg", "Hello from Container!"]
```
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o server .
EXPOSE 5000

# 항상 실행되는 진입점 (고정)

## 2. 핵심 개념 — ENTRYPOINT와 CMD

- **ENTRYPOINT** = 컨테이너가 실행할 **고정된 명령**. `docker run`의 인자로는 바뀌지 않습니다(바꾸려면 `--entrypoint` 플래그가 필요).
- **CMD** = ENTRYPOINT에 넘길 **기본 인자**. `docker run <이미지> <인자...>` 처럼 이미지 뒤에 인자를 주면 CMD가 통째로 그 인자로 **대체**됩니다.
- 둘을 조합하면 실제 실행은 `ENTRYPOINT + (CMD 또는 사용자가 준 인자)` 가 됩니다.

| 명령 | 실제 실행 |
|------|-----------|
| `docker run counter:2.0` | `./server -msg "Hello from Container!"` |
| `docker run counter:2.0 -msg "Hello (prod)"` | `./server -msg "Hello (prod)"` |
| `docker run --entrypoint sh counter:2.0 -c "..."` | `sh -c "..."` (진입점 자체를 교체) |

## 3. exec form을 쓰는 이유

`["./server"]` 같은 **exec form**(JSON 배열)은 셸을 거치지 않고 바이너리를 PID 1로 직접 실행합니다. 반면 `ENTRYPOINT ./server` 같은 **shell form**은 `/bin/sh -c`로 감싸져 실행되어, 시그널(SIGTERM 등)이 프로세스에 제대로 전달되지 않습니다. `docker stop`이 깔끔히 동작하려면 exec form을 쓰세요.

## 4. 언제 무엇을

- 컨테이너가 **하나의 고정된 프로그램**으로 동작해야 하면 → 그 프로그램을 `ENTRYPOINT`로.
- 그 프로그램에 **바꿔 끼울 기본값**이 있으면 → `CMD`로.
- "이미지가 곧 하나의 실행 파일처럼" 동작하게 만드는 패턴입니다.
