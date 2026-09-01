FROM docker.io/storjlabs/satellite:4bf56a418-v1.102.1-rc-go1.21.3
WORKDIR /app

COPY cmd/metasearch/entrypoint /entrypoint

EXPOSE 6666

ENV GOARCH=amd64