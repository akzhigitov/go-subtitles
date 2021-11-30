FROM golang:1.17.3 AS build
ENV GO111MODULE=on
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o /app/out/go-subtitles .
FROM scratch AS bin
COPY --from=build /app/out/go-subtitles /go-subtitles
COPY public /public/
ENTRYPOINT ["/go-subtitles"]