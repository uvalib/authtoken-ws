if [ -z "$DOCKER_HOST" ]; then
   echo "ERROR: no DOCKER_HOST defined"
   exit 1
fi

# set the definitions
INSTANCE=authtoken-ws
NAMESPACE=uvadave

docker run -ti -p 8180:8080 $NAMESPACE/$INSTANCE:latest /bin/bash
