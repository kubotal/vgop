# Build the manager binary
FROM golang:1.22 AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/main.go cmd/main.go
COPY api/ api/
COPY internal/ internal/

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o vgop cmd/main.go


FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive
# Install some misc tools and base python3 stuff
RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get install -y tini util-linux lvm2 && \
    rm -rf /var/lib/apt/lists/* &&  \
    rm -rf /var/cache/apt/*

WORKDIR /
COPY --from=builder /workspace/vgop .
USER 0:0

ENTRYPOINT ["/vgop"]
