FROM alpine:3.10

# update the packages
RUN apk update && apk upgrade && apk add bash tzdata ca-certificates && rm -fr /var/cache/apk/*

# Create the run user and group
RUN addgroup --gid 18570 sse && adduser --uid 1984 docker -G sse -D

# set the timezone appropriatly
ENV TZ=UTC
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Specify home 
ENV APP_HOME /authtoken-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin $APP_HOME/assets
RUN chown -R docker $APP_HOME && chgrp -R sse $APP_HOME

# Add the RDS certificates
COPY data/rds-combined-ca-bundle.pem /etc/ssl/certs

# Specify the user
USER docker

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary assets
COPY data/container_bash_profile /home/docker/.profile
COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY bin/authtoken-ws.linux $APP_HOME/bin/authtoken-ws
COPY assets/* $APP_HOME/assets/

# Add the build tag
COPY buildtag.* $APP_HOME/
