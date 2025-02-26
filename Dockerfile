FROM alpine:3

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

