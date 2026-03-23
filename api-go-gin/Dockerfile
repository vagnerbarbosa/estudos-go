FROM postgres:15-alpine

ENV POSTGRES_PASSWORD=1234
ENV POSTGRES_DB=api_go_gin

COPY init.sql /docker-entrypoint-initdb.d/