
FROM golang:1.20 as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

FROM alpine:3.17

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata libc6-compat
ENV TZ=Asia/Bangkok

COPY --from=builder /app/main .

RUN apk add --no-cache libcap && \
    setcap 'cap_net_bind_service=+ep' /app/main && \
    apk del libcap && \
    adduser -D -u 1000 nonroot
USER nonroot

ENTRYPOINT ["./main"]
