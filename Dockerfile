FROM golang:1.14.15

WORKDIR /app

COPY . ./

RUN apt-get update && \
  apt-get install -y --no-install-recommends \
    g++ \
    git \
    libssl-dev \
    swig

CMD [ "./gen_lib" ]
