FROM golang:1.15

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /go/src/gin-mysql-api   

RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

COPY server/ ./

RUN go build -o gma . 

EXPOSE 8080

ENTRYPOINT ["./gma"]
