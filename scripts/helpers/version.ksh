#
#
#

# ensure we have and endpoint
if [ -z "$TOKENAUTH_URL" ]; then
   echo "ERROR: TOKENAUTH_URL is not defined"
   exit 1
fi

# issue the command
echo "$TOKENAUTH_URL"
curl $TOKENAUTH_URL/version

exit 0

#
# end of file
#
