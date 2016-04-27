if [ -z "$DOCKER_HOST" ]; then
   echo "ERROR: no DOCKER_HOST defined"
   exit 1
fi

# set the definitions
INSTANCE=authtoken-ws
NAMESPACE=uvadave

# stop the running instance
docker stop $INSTANCE

# remove the instance
docker rm $INSTANCE

# remove the previously tagged version
docker rmi $NAMESPACE/$INSTANCE:current  

# tag the latest as the current
docker tag -f $NAMESPACE/$INSTANCE:latest $NAMESPACE/$INSTANCE:current

docker run -d -p 8200:8080 -e DBPASSWD=$DBPASSWD --name $INSTANCE $NAMESPACE/$INSTANCE:latest

# return status
exit $?
