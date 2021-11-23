FROM golang:1.17
WORKDIR /root

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    locales \
    && apt-get -y clean \
    && rm -rf /var/lib/apt/lists/*

RUN sed -i -e "s/# ja_JP.UTF-8 UTF-8/ja_JP.UTF-8 UTF-8/" /etc/locale.gen

RUN locale-gen ja_JP.utf8
ENV LANG="ja_JP.utf8" \
    LANGUAGE="ja_JP.utf8" \
    LC_ALL="ja_JP.utf8"

RUN go get github.com/cespare/reflex