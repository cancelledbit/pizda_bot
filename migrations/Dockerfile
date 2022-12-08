FROM gomicro/goose

ADD *.sql /migrations/
ADD entrypoint.sh /migrations/
RUN chmod +x /migrations/entrypoint.sh

ENTRYPOINT ["/migrations/entrypoint.sh"]