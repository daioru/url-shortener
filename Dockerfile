# Builder

FROM golang:1.23.6-alpine AS builder

ARG GITHUB_PATH=github.com/daioru/url-shortener

COPY Makefile /home/${GITHUB_PATH}/Makefile
COPY go.mod /home/${GITHUB_PATH}/go.mod
COPY go.sum /home/${GITHUB_PATH}/go.sum

WORKDIR /home/${GITHUB_PATH}

COPY . /home/${GITHUB_PATH}

RUN apk add --no-cache make
RUN make build


# Server

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

ARG GITHUB_PATH=github.com/daioru/url-shortener

COPY --from=builder /home/${GITHUB_PATH}/bin/url-shortener .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
# COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root url-shortener

EXPOSE 8080

CMD ["./url-shortener"]