# build stage
FROM golang

# # librdkafka Build from source
WORKDIR /tmp
RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR /tmp/librdkafka
RUN ./configure --prefix /usr
RUN make
RUN make install

RUN go get -insecure github.com/mixi-gaminh/core-framework