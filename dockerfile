FROM ubuntu
WORKDIR /app

ENV HOME="/" \
    OS_ARCH="amd64" \
    OS_FLAVOUR="ubuntu" \
    OS_NAME="linux"

LABEL org.opencontainers.image.source="https://github.com/hongque-pro/infra-tracing-es-apm-sink"

COPY output .

ENV PATH="/app:$PATH"

CMD [ "tracing-sink" ]
