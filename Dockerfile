FROM amd64/debian:stable-slim

# set noninteractive mode so those terrible debconf
# dialog initialization errors go away
RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections

# do these steps first, so docker can cache these image layers independent
# of the entrypoint and webproxy. that way, if the webproxy build changes,
# we don't need to run these again to rebuild the docker container.
RUN apt-get update

COPY entrypoint /app/entrypoint

# copy go binary
COPY longform /app/longform

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint"]

