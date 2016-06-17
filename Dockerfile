FROM alpine:3.4

RUN apk --update add bash

# Create the run user and group
RUN addgroup webservice && adduser webservice -G webservice -D

# Specify home 
ENV APP_HOME /authtoken-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin
RUN chown -R webservice $APP_HOME && chgrp -R webservice $APP_HOME

# Specify the user
USER webservice

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary assets
COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY bin/authtoken-ws.alpine $APP_HOME/bin/authtoken-ws
