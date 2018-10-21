FROM golang:1.11-alpine AS gobuild
RUN apk add curl make
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/gitlab.com/rafadc/ddns-update/

COPY . ./

RUN dep ensure
RUN make build

FROM alpine:3.6

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

RUN mkdir /app
COPY --from=gobuild /go/src/gitlab.com/rafadc/ddns-update/target/ddns-update /app
CMD ["/app/ddns-update"]
