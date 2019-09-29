FROM alpine:latest
WORKDIR /go/src/hello
COPY hello /go/src/hello

# RUN CGO_ENABLED=0 GOOS=linux go build -o hello

EXPOSE 8000
CMD ["./hello"]