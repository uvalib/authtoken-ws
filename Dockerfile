FROM centos:7

ENV APP_HOME /authtoken-ws
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin

EXPOSE 8080
CMD scripts/entry.sh
WORKDIR $APP_HOME

COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY bin/authtoken-ws.linux $APP_HOME/bin/authtoken-ws
