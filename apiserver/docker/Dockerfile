# ベースとなるDockerイメージ指定
FROM golang:latest
#dockerize初期設定
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && apt-get update \
    && apt-get install vim \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/gin-contrib/cors \
    && go get github.com/gin-gonic/gin

# コンテナ内に作業ディレクトリを作成
RUN mkdir -p /go/src/github.com/kamio06/mystery_land/apiserver
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/github.com/kamio06/mystery_land/apiserver
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/kamio06/mystery_land/apiserver
RUN go get -u github.com/gin-gonic/gin

EXPOSE 8800
