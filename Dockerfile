
FROM golang:1.19-alpine as builder

ADD ./ /go-charts

RUN apk update \
  && apk add docker git gcc make \
  && cd /go-charts \
  && make build

FROM alpine 

EXPOSE 7001

COPY --from=builder /go-charts/go-charts /usr/local/bin/go-charts
COPY --from=builder /go-charts/entrypoint.sh /entrypoint.sh


CMD ["go-charts"]

ENTRYPOINT ["/entrypoint.sh"]

HEALTHCHECK --timeout=10s --interval=10s CMD [ "wget", "http://127.0.0.1:7001/ping", "-q", "-O", "-"]