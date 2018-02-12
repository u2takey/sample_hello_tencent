FROM alpine:3.6
ADD make/release/linux/amd64/hello  /hello
EXPOSE 8808
ENTRYPOINT ["/hello"]