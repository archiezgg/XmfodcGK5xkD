#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/library-mgtm /library-mgtm
LABEL Name=library-mgmt Version=0.0.1
EXPOSE 8080
CMD ["/library-mgtm"]