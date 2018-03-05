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

LT=../../bin/bombardier
if [ ! -f "$LT" ]; then
   echo "ERROR: Bombardier is not available"
   exit 1
fi

# set the test parameters
endpoint=$TOKENAUTH_URL
concurrent=10
count=10000
url=authorize/loadtest/what/$API_TOKEN

CMD="$LT -c $concurrent -n $count -l $endpoint/$url"
echo "Host = $TOKENAUTH_URL, count = $count, concurrency = $concurrent"
echo $CMD
$CMD
exit $?

#
# end of file
#
