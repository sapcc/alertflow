FROM keppel.eu-de-1.cloud.sap/ccloud-dockerhub-mirror/library/golang:1.19-alpine as builder
RUN apk add --no-cache make git

WORKDIR /go/src/github.com/sapcc/alertflow

COPY go.mod go.sum ./
RUN go mod download

COPY pkg/ pkg/
COPY main.go main.go
COPY Makefile Makefile
RUN make all

FROM keppel.eu-de-1.cloud.sap/ccloud-dockerhub-mirror/library/alpine:latest
LABEL source_repository="https://github.com/sapcc/alertflow"
LABEL org.opencontainers.image.authors="Bassel Zeidan <bassel.zeidan@sap.com>"

COPY --from=builder /go/src/github.com/sapcc/alertflow/bin/alertflow /usr/local/bin/
CMD ["alertflow"]
