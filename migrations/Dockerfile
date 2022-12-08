FROM gomicro/goose

ADD *.sql /migrations/
ADD entrypoint.sh /migrations/
RUN apk update && apk add curl
RUN chmod +x /migrations/entrypoint.sh

ENTRYPOINT ["/migrations/entrypoint.sh"]