
###################################
#Build stage
FROM golang:1.16-buster AS build-env

ARG GITEA_VERSION
ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata $TAGS"

#Build deps
#RUN apk --no-cache add build-base git
RUN apt-get install git -y

# Install wkhtmltopdf
RUN apt-get update -y
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb
RUN dpkg -i wkhtmltox_0.12.6-1.buster_amd64.deb || true
RUN apt-get install -f -y
RUN ldconfig
RUN rm wkhtmltox_0.12.6-1.buster_amd64.deb

#Setup repo
COPY . ${GOPATH}/src/github.com/masudur-rahman/pdf-generator
WORKDIR ${GOPATH}/src/github.com/masudur-rahman/pdf-generator
RUN go mod init && go mod tidy && go mod vendor

#RUN export XDG_RUNTIME_DIR=/tmp

#Checkout version if set
RUN CGO_ENABLED=1 GOOS=linux go build -o pdf-generator .

FROM debian:buster-slim
LABEL maintainer="maintainers@gitea.io"

EXPOSE 22 3000

RUN apt-get update -y
RUN apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    git \
    gettext \
    s6 \
    sqlite \
    tzdata

RUN addgroup \
    --system --gid 1000 \
    git && \
  adduser \
    --system \
    --home /data/git \
    --shell /bin/bash \
    --uid 1000 \
    --gid 1000 \
    git

RUN echo "git:$(dd if=/dev/urandom bs=24 count=1 status=none | base64)" | chpasswd


RUN apt search wget
RUN apt-get install wget -y

# Install wkhtmltopdf
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb
RUN dpkg -i wkhtmltox_0.12.6-1.buster_amd64.deb || true
RUN apt-get install -f -y
RUN ldconfig
RUN rm wkhtmltox_0.12.6-1.buster_amd64.deb

COPY --from=build-env /go/src/github.com/masudur-rahman/pdf-generator/pdf-generator /app/pdf-generator/pdf-generator
COPY --from=build-env /go/src/github.com/masudur-rahman/pdf-generator/templates/simple.html /templates/simple.html

ENTRYPOINT ["/app/pdf-generator/pdf-generator"]
