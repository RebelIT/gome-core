FROM golang:1.15.0-buster

WORKDIR "/gome-core"
ADD entrypoint.sh /
RUN ["chmod", "+x", "/entrypoint.sh"]
ENTRYPOINT ["/entrypoint.sh"]