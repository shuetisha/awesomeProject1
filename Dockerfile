FROM alpine:latest
ENTRYPOINT ["/usr/bin/resmo-db-mapper"]
COPY pscale /usr/bin