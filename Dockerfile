FROM golang:1.21-alpine

RUN apk update && apk add git

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /go/src/app

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN go install github.com/volatiletech/sqlboiler@latest
RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air","-c",".air.toml"]