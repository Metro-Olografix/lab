FROM ubuntu

# install stuff
RUN apt-get update \
    && apt-get install -y \
    sudo \
    curl \
    wget \
    git \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# install golang
ARG GO_VERSION
ENV GO_VERSION=${GO_VERSION}

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN mkdir /usr/local/go && tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# run as
ARG USER=coder
RUN useradd --no-create-home --shell /bin/bash ${USER}
USER ${USER}

WORKDIR /home/${USER}
