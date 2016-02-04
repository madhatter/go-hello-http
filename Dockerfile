FROM scratch
MAINTAINER Arvid Warnecke 'madhatter@nostalgix.org'

EXPOSE 8080

WORKDIR /app
COPY hello /app/

ENTRYPOINT ["./hello"]
