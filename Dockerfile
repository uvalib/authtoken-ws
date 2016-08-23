FROM alpine:3.4

# temp cos the CDN is down
RUN sed -i -e 's/dl-cdn/dl-4/g' /etc/apk/repositories

# Add bash and timezone data cos we dont get them by default
RUN apk --update add bash tzdata

# Create the run user and group
RUN addgroup webservice && adduser webservice -G webservice -D

# set the timezone appropriatly
ENV TZ=EST5EDT
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

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
COPY data/container_bash_profile /home/webservice/.profile
COPY bin/authtoken-ws.linux $APP_HOME/bin/authtoken-ws

# Add the build tag
COPY buildtag.* $APP_HOME/
