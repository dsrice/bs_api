FROM golang:1.21-alpine

RUN apk update && apk add git

WORKDIR /go/src/app

RUN go install github.com/volatiletech/sqlboiler@latest
RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest