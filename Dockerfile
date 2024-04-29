FROM rust:slim-buster AS builder

WORKDIR /app
COPY . .
RUN cargo build -r
RUN cp ./target/release/jester /

FROM alpine:latest AS final

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    appuser
COPY --from=builder /jester /usr/local/bin

RUN chown appuser /usr/local/bin/jester
USER appuser

WORKDIR /jester
ENTRYPOINT ["jester"]