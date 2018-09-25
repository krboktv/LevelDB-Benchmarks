FROM golang:latest

WORKDIR /app
COPY . .

RUN git clone https://github.com/google/leveldb.git

WORKDIR /app/leveldb

RUN mkdir -p build

WORKDIR /app/leveldb/build

ADD https://cmake.org/files/v3.12/cmake-3.12.2-Linux-x86_64.sh /cmake-3.7.2-Linux-x86_64.sh
RUN mkdir /opt/cmake
RUN sh /cmake-3.7.2-Linux-x86_64.sh --prefix=/opt/cmake --skip-license
RUN ln -s /opt/cmake/bin/cmake /usr/local/bin/cmake

RUN cmake -DCMAKE_BUILD_TYPE=Release ..
RUN cmake --build .

WORKDIR /app

RUN go get github.com/syndtr/goleveldb/leveldb

ENV SECONDS=10

RUN go build -o main .

CMD ["/app/main"]