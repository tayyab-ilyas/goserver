FROM debian:stable-slim

COPY goserver /bin/goserver

ENV PORT 8000

CMD ["/bin/goserver"]