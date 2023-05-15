FROM alpine:latest
COPY awesomeProject1 /awesomeProject1/awesomeProject1
WORKDIR "/awesomeProject1"
ENTRYPOINT ["/awesomeProject1/awesomeProject1"]