FROM node:12.7.0-alpine

WORKDIR /myapp

# install golang
RUN apk add --update --no-cache vim git make musl-dev go curl && \
    export GOPATH= && \
    export PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH && \
    export GOBIN=$GOROOT/bin && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin && \
    export GO111MODULE=on

# install jq
RUN curl -o /usr/local/bin/jq -L https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 && chmod +x /usr/local/bin/jq
# install go-swagger
RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') && \
    curl -o /usr/local/bin/swagger -L'#' "$download_url" && \
    chmod +x /usr/local/bin/swagger