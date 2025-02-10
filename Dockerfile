FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY * ./
RUN go mod download

ENV CGO_ENABLED 0
ENV GOOS ${TARGETOS}
ENV GOARCH ${TARGETARCH}

RUN go build -ldflags="-s -w" -o bin/hook-dumper

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch as run

WORKDIR /app
COPY --from=builder /app/bin/hook-dumper /hook-dumper

ENTRYPOINT ["/hook-dumper"]
EXPOSE 8080
