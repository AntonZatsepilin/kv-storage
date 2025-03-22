FROM tarantool/tarantool:2.11.0

RUN apk add --no-cache --virtual .build-deps build-base git && \
    tarantoolctl rocks install http && \
    apk del .build-deps

WORKDIR /opt/tarantool
COPY init.lua healthcheck.lua ./

CMD ["tarantool", "init.lua"]