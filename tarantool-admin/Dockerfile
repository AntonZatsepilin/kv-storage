FROM tarantool/tarantool:latest

# Install HTTP module
RUN tarantoolctl rocks install http

# Copy initialization scripts
COPY init.lua /opt/tarantool/
COPY healthcheck.lua /opt/tarantool/

EXPOSE 3301 8081
CMD ["tarantool", "/opt/tarantool/init.lua"]