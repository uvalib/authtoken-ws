if [ -z "$DOCKER_HOST" ]; then
   echo "ERROR: no DOCKER_HOST defined"
   exit 1
fi

# set the definitions
INSTANCE=authtoken-ws

CID=$(docker ps -f name=$INSTANCE|grep -v jetty|tail -1|awk '{print $1}')
if [ -n "$CID" ]; then
   docker exec -it $CID /bin/bash -l
else
   echo "No running container for $INSTANCE"
fi

