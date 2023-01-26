FROM golang:1.19 AS build

RUN useradd -u 10001 immunity

WORKDIR /build/
COPY . /build/

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor

FROM busybox AS package

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /build/immunity .
COPY benthos/benthos.yaml /benthos.yaml

USER immunity

EXPOSE 4195
EXPOSE 8181

ENTRYPOINT ["/immunity"]