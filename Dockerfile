FROM golang:1.13 as build
ADD . /go/src/github.com/wangcanfengxs/mock-server
WORKDIR  /go/src/github.com/wangcanfengxs/mock-server
RUN CGO_ENABLED=0 GOOS=linux go build -o /mock-server -ldflags '-s -w' main.go
RUN chmod +x /mock-server

FROM alpine:3.7

COPY --from=build /mock-server /mock-server
RUN ln -s /mock-server /usr/bin/mock-server

ENTRYPOINT [ "/mock-server" ]