FROM alpine:3

COPY files/github-cli /usr/local/bin/

RUN set -ex \
  && apk add --no-cache ca-certificates

ENTRYPOINT ["github-cli"]
CMD ["--help"]
