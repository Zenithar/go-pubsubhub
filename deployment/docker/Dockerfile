# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

## -------------------------------------------------------------------------------------------------

FROM golang:1.12 as builder

RUN set -eux; \
    apt-get update -y && \
    apt-get install -y apt-utils upx zip unzip;

# Create a non-root privilege account to build
RUN adduser --disabled-password --gecos "" -u 1000 golang && \
    mkdir -p $GOPATH/src/workspace && \
    chown -R golang:golang $GOPATH/src/workspace;

# Force go modules
ENV GO111MODULE=on

WORKDIR $GOPATH/src/workspace

# Prepare an unprivilegied user for run
RUN set -eux; \
    echo 'nobody:x:65534:65534:nobody:/:' > /tmp/passwd && \
    echo 'nobody:x:65534:' > /tmp/group && \
    mkdir /tmp/.config && \
    chown 65534:65534 /tmp/.config

# Drop privileges to build
USER golang
COPY --chown=golang:golang . .

# Install dependencies
RUN set -eux; \
    go run mage.go -d tools && \
    go run mage.go go:deps

# Build final target
RUN set -eux; \
    go run mage.go

# Compress binaries
RUN set -eux; \
    upx -9 bin/* && \
    chmod +x bin/*

## -------------------------------------------------------------------------------------------------

FROM gcr.io/distroless/static:latest

# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

# Metadata
LABEL \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="PubSubHub" \
    org.label-schema.description="Lightweight messaging system based on gRPC." \
    org.label-schema.url="https://go.zenithar.org/pubsubhub" \
    org.label-schema.vcs-url="https://github.com/Zenithar/go-pubsubhub.git" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vendor="Thibault NORMAND" \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0" \
    org.zenithar.licence="MIT"

COPY --from=builder /go/src/workspace/bin/pubsubhub /usr/bin/pubsubhub
COPY --from=builder /tmp/group /tmp/passwd /etc/
COPY --from=builder --chown=65534:65534 /tmp/.config /

USER nobody:nobody
WORKDIR /

ENTRYPOINT [ "/usr/bin/pubsubhub" ]
CMD ["--help"]