#
# basic load test
#

if [ -z "$TOKENAUTH_URL" ]; then
   echo "ERROR: TOKENAUTH_URL is not defined"
   exit 1
fi

if [ -z "$API_TOKEN" ]; then
   echo "ERROR: API_TOKEN is not defined"
   exit 1
fi

AB=$(which ab | head -1)
if [ -z "$AB" ]; then
   echo "ERROR: Apache Bench is not available"
   exit 1
fi

# set the test parameters (note time and count are alternatives, test will not be both)
endpoint=$TOKENAUTH_URL
concurrent=5
time=15
count=5000
url=authorize/loadtest/what/$API_TOKEN

CMD="$AB -t $time -c $concurrent -n $count $endpoint/$url"
echo $CMD
$CMD
exit $?

#
# end of file
#
