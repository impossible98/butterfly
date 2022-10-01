# one
FROM impossible98/go
WORKDIR /src/
COPY ./ ./
RUN make install \
    && make build
# two
FROM debian:10-slim
WORKDIR /
COPY --from=0 /src/dist/butterfly ./
EXPOSE 7887
ENTRYPOINT ["/butterfly", "serve"]
