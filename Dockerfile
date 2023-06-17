FROM lscr.io/linuxserver/raneto:latest
LABEL authors="NekoLuka"

COPY --from=golang:alpine3.18 /usr/local/go /usr/local/go
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /repowiki
COPY ./ /repowiki

ENTRYPOINT go run .