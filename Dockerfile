FROM golang:1.11-alpine AS gobuild
RUN apk add build-base make git

WORKDIR /ddns-update/

COPY . ./

RUN go mod download
RUN make build

FROM alpine:3.6

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

RUN mkdir /app
COPY --from=gobuild /ddns-update/target/ddns-update /app
CMD ["/app/ddns-update"]
