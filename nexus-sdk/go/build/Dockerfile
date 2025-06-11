FROM --platform=${BUILDPLATFORM} golang:1.24-bookworm AS build_stage

WORKDIR /go/src

COPY ./go .
COPY ./_distfiles_/SDK/include/libdfi.h /go/src/pkg/scanner/nexus/include/libdfi.h
COPY ./_distfiles_/SDK/lib/linux/arm64/libdfi.so /go/src/pkg/scanner/nexus/lib/arm64/libdfi.so
COPY ./_distfiles_/SDK/lib/linux/x64/libdfi.so /go/src/pkg/scanner/nexus/lib/amd64/libdfi.so

RUN set -e && \
    go mod tidy && \
    go build -o s1scanner ./cmd/s1scanner

FROM --platform=${BUILDPLATFORM} busybox:glibc

ARG TARGETARCH
ARG TARGETOS

WORKDIR /

COPY --from=build_stage /go/src/s1scanner /bin/s1scanner
COPY ./_distfiles_/SDK/lib/linux/arm64/libdfi.so /lib/libnexus-arm64.unstripped.so
COPY ./_distfiles_/SDK/lib/linux/x64/libdfi.so /lib/libnexus-amd64.unstripped.so
COPY ./sample-files /opt/s1scanner/sample-files/

RUN set -e && \
    ln -sf libnexus-${TARGETARCH}.unstripped.so /lib/libnexus.unstripped.so && \
    chmod +x /bin/s1scanner

ENTRYPOINT ["/bin/s1scanner"]
