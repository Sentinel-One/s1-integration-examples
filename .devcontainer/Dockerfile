FROM mcr.microsoft.com/devcontainers/base:bookworm

ARG GO_VERSION=1.23.9
ARG NVM_VERSION=0.40.3

COPY files /

# base setup
RUN set -xe \
    && apt update \
    && apt dist-upgrade -y \
    && cp /root/.bashrc /home/vscode/.bashrc

# install Python tools
RUN set -xe \
    && apt install -y python3 python3-pip

# install Go tools
RUN set -xe \ 
    && cd /tmp \
    && if [ "$(uname -m)" = "aarch64" ]; then ARCH=arm64; else ARCH=amd64; fi \
    && wget https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz \
    && tar -zxvf go${GO_VERSION}.linux-${ARCH}.tar.gz -C /usr/local \
    && export GOPATH=/usr/local/go \
    && /usr/local/go/bin/go install golang.org/x/pkgsite/cmd/pkgsite@latest \
    && rm -rf /tmp/*

# install NodeJS tools
RUN set -xe \
    && cd /tmp \
    && wget https://raw.githubusercontent.com/nvm-sh/nvm/v${NVM_VERSION}/install.sh \
    && chmod +x install.sh \
    && mkdir /home/vscode/.nvm \
    && chown -hR vscode /home/vscode \
    && sudo -i -u vscode /tmp/install.sh \
    && sudo -i -u vscode -- nvm install --lts
