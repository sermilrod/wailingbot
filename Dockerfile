FROM golang:1.9.2-alpine3.6
ENV APP_PATH /go/src/github.com/sermilrod/wailingbot
RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH
ADD . $APP_PATH
RUN apk add --update git
RUN rm -rf vendor
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build

FROM alpine:3.6
COPY --from=0 /go/src/github.com/sermilrod/wailingbot/wailingbot .
RUN mkdir migrations
COPY --from=0 /go/src/github.com/sermilrod/wailingbot/migrations migrations/
