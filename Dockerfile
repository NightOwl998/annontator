# syntax=docker/dockerfile:1
# build stage
FROM golang:1.19-alpine 
RUN apk add --no-cache git

WORKDIR  /annontator
#RUN make vendor
#ADD . /go/src/github.com/martonsereg/scheduler
#RUN go mod d
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
RUN go build -o ./ .
RUN ls -l
#RUN go build -o /annontator
CMD [ "./annontator" ]

