# Dockerfile
FROM golang:1.23.3-bookworm as builder
RUN mkdir /go-http-server
COPY main.go /go-http-server/main.go
WORKDIR /go-http-server
RUN go mod init go-http-server
RUN CGO_ENABLED=0 go build -o server

# stage 2
FROM gcr.io/distroless/static-debian12
COPY --from=builder /go-http-server/server / 
ENTRYPOINT [ "/server" ]

# docker buildx build --platform linux/amd64 -t go-http-server:0.1 .
# docker run -d --name  go-http-server  -p 8888:8080   go-http-server:0.1 
# curl localhost:8888

