FROM alpine:latest
ENTRYPOINT ["/usr/bin/resmo-db-mapper"]
COPY resmo-db-mapper /usr/bin