if [ -z "$DOCKER_HOST" ]; then
   echo "ERROR: no DOCKER_HOST defined"
   exit 1
fi

echo "*****************************************"
echo "building on $DOCKER_HOST"
echo "*****************************************"

# set the definitions
INSTANCE=authtoken-ws
NAMESPACE=uvadave

# pull base image to ensure we have the latest
docker pull alpine:3.10

# build the image
docker build -t $NAMESPACE/$INSTANCE .

# return status
exit $?
