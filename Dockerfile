FROM golang:1.21-alpine

RUN apk update && apk add git

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /go/src/app

RUN go install github.com/volatiletech/sqlboiler@latest
RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]