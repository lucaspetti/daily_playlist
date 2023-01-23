FROM golang:1.16-alpine

ADD . /go/src/daily_playlist

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /daily_playlist

ENTRYPOINT [ "./daily_playlist" ]
