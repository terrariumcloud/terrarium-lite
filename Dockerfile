FROM golang:1.17.1 as build
ENV CGO_ENABLED=0 GOOS=linux GARCH=amd64
WORKDIR /workspace
COPY . /workspace
RUN go build -o terrarium
RUN apt-get update && \
    apt-get install -y ca-certificates

FROM busybox
COPY --from=build /workspace/terrarium /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "/terrarium" ]
CMD ["serve"]