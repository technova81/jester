FROM rust:alpine AS builder

WORKDIR /app
COPY . .
RUN cargo build -r
RUN cp ./target/release/athena-bot /

FROM alpine:latest AS final

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    appuser
COPY --from=builder /athena-bot /usr/local/bin

RUN chown appuser /usr/local/bin/athena-bot
COPY --from=builder /app/config /opt/athena-bot/config

RUN chown -R appuser /opt/athena-bot
USER appuser

WORKDIR /opt/athena-bot
ENTRYPOINT ["athena-bot"]