if [ -z "$DBPASSWD" ]; then
   echo "ERROR: DBPASSWD must be defined"
   exit 1
fi

bin/authtoken-ws.darwin --dbpassword $DBPASSWD
