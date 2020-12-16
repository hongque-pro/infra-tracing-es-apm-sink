FROM ubuntu
WORKDIR /app

ENV HOME="/" \
    OS_ARCH="amd64" \
    OS_FLAVOUR="ubuntu" \
    OS_NAME="linux"

COPY output .

ENV PATH="/app:$PATH"

CMD [ "tracing-sink" ]
