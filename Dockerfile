FROM golang:1.20 as builder

WORKDIR /app

ENV GOPROXY=https://goproxy.io,direct

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./main.go

FROM scratch

COPY --from=builder /app/myapp /myapp

COPY --from=builder /app/config /config

EXPOSE 8088

ENTRYPOINT ["/myapp"]
