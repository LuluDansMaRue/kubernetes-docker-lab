FROM rust:1.38-stretch

WORKDIR /usr/src/chat
COPY chat .

RUN cargo install --path .

CMD ["chat"]