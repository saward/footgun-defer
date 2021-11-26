FROM postgres:14

# Broke CI tests: https://github.com/docker-library/postgres/issues/681
ENV POSTGRES_HOST_AUTH_METHOD trust

HEALTHCHECK --interval=30s --timeout=5s --retries=5 \
    CMD pg_isready -U postgres

COPY migrations/*.sql /docker-entrypoint-initdb.d/
