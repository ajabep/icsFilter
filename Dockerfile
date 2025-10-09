FROM alpine:3@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412

LABEL org.opencontainers.image.authors="ajabep"
LABEL org.opencontainers.image.url="https://github.com/ajabep/icsFilter"
LABEL org.opencontainers.image.source="https://github.com/ajabep/icsFilter"
LABEL org.opencontainers.image.licenses="Unlicense"

WORKDIR /
EXPOSE 8080/tcp
RUN apk add --no-cache curl=8.12.1-r0 tzdata=2025a-r0
HEALTHCHECK --interval=10s CMD curl http://127.0.0.1:8080
ENTRYPOINT ["/icsFilter"]

COPY icsFilter /

USER nobody

