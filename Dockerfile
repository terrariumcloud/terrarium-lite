FROM golang:1.17.1 as build
ENV CGO_ENABLED=0 GOOS=linux GARCH=amd64
COPY . /workspace
RUN cd /workspace && \
    go build -o terrarium

FROM scratch
COPY --from=build /workspace/terrarium /
ENTRYPOINT [ "/terrarium" ]
CMD ["serve"]