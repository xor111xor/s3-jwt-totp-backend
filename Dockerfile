FROM alpine

ENV PORT 9090

RUN mkdir /app
ADD  s3-jwt-totp-backend /app
ADD  data/example/config/storage-server.toml /app/storage-server.toml
COPY  templates/ /app/templates/
WORKDIR /app

EXPOSE $PORT

ENTRYPOINT /app/s3-jwt-totp-backend --config /app/storage-server.toml

