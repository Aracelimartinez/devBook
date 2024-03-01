FROM postgres:14.10

COPY *.sql /docker-entrypoint-initdb.d/

CMD ["postgres", "-p", "5432"]
