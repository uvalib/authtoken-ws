if [ -z "$DOCKER_HOST" ]; then
   echo "ERROR: no DOCKER_HOST defined"
   exit 1
fi

# set the definitions
INSTANCE=authtoken-ws
NAMESPACE=uvadave

DB_ENV="-e DBHOST=$DBHOST -e DBNAME=$DBNAME -e DBUSER=$DBUSER -e DBPASSWD=$DBPASSWD"

# stop the running instance
docker stop $INSTANCE

# remove the instance
docker rm $INSTANCE

# remove the previously tagged version
docker rmi $NAMESPACE/$INSTANCE:current  

# tag the latest as the current
docker tag -f $NAMESPACE/$INSTANCE:latest $NAMESPACE/$INSTANCE:current

docker run -d -p 8200:8080 $DB_ENV --name $INSTANCE $NAMESPACE/$INSTANCE:latest

# return status
exit $?
