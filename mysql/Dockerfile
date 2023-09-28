FROM mariadb:10.11

# RUN apt-get update 
# RUN apt-get -y install git wget curl


# ARG GO_VER=1.21.1
# ARG GO_BIN=/usr/local/go/bin
ARG SODA_BIN=./pop_6.1.1_linux_amd64/soda

# install go
# RUN wget https://go.dev/dl/go${GO_VER}.linux-amd64.tar.gz && \
#     tar -C /usr/local -xzf go${GO_VER}.linux-amd64.tar.gz  && \
#     rm *.tar.gz

# ENV PATH=$PATH:${GO_BIN}

# # install soda for sql migration
# RUN go install github.com/gobuffalo/pop/v6/soda@latest && \
#     soda

COPY ${SODA_BIN} /usr/local/bin/soda

RUN mkdir /workspace/
WORKDIR /workspace

EXPOSE 3306


