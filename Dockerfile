FROM golang:alpine3.18

LABEL authors="NekoLuka"

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /repowiki
COPY ./ /repowiki
RUN mkdir bin && go build -o bin/repowiki .

CMD /repowiki/bin/repowiki