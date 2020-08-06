FROM golang:1.14-alpine AS build

WORKDIR /socks5-server/
COPY . /socks5-server

RUN CGO_ENABLED=0 go build -ldflags -v -o /bin/socks5-server cmd/main.go

FROM scratch
COPY --from=build /bin/socks5-server /bin/socks5-server
ENTRYPOINT ["/bin/socks5-server"]