if [ -z "$DBPASSWORD" ]; then
   echo "ERROR: DBPASSWORD must be defined"
   exit 1
fi

bin/authtoken-ws.darwin --dbpassword $DBPASSWORD
