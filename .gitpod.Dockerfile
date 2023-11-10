FROM public.ecr.aws/z2f7n8a1/couchbase-da-containers:couchbase-neo

RUN echo "* soft nproc 20000\n"\
    "* hard nproc 20000\n"\
    "* soft nofile 200000\n"\
    "* hard nofile 200000\n" >> /etc/security/limits.conf

RUN apt-get -qq update && \
    apt-get install -yq sudo git wget

RUN cd /tmp && \
    wget --no-check-certificate https://golang.org/dl/go1.19.1.linux-amd64.tar.gz && \
    tar -xzf go1.19.1.linux-amd64.tar.gz && \
    mv go /usr/local/go && \
    echo 'export PATH=$PATH:/usr/local/go/bin && \
    export GOPATH=/var/www/go && \
    export PATH=$PATH:$GOPATH/bin' >> /etc/profile && \
    . /etc/profile && \
    mkdir -p $GOPATH/github.com/couchbase-examples/golang-quickstart && \
    mkdir -p $GOPATH/bin && \
    mkdir -p $GOPATH/pkg && \
    go version



COPY startcb.sh /opt/couchbase/bin/startcb.sh
USER gitpod
ENV PATH="/usr/local/go/bin:$PATH"
