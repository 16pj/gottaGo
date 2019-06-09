FROM ubuntu:16.04

COPY main.go /main.go

RUN apt-get update && \
    apt-get install -y golang

RUN go build /main.go

ENTRYPOINT ./main
    
