# set blank options variables
DBSECURE_OPT=""
DBHOST_OPT=""
DBNAME_OPT=""
DBUSER_OPT=""
DBPASSWD_OPT=""
DEBUG_OPT=""

# secure database access
if [ -n "$DBSECURE" ]; then
   DBSECURE_OPT="--dbsecure=$DBSECURE"
fi

# database host
if [ -n "$DBHOST" ]; then
   DBHOST_OPT="--dbhost $DBHOST"
fi

# database name
if [ -n "$DBNAME" ]; then
   DBNAME_OPT="--dbname $DBNAME"
fi

# database user
if [ -n "$DBUSER" ]; then
   DBUSER_OPT="--dbuser $DBUSER"
fi

# database password
if [ -n "$DBPASSWD" ]; then
   DBPASSWD_OPT="--dbpassword $DBPASSWD"
fi

# service debugging
if [ -n "$AUTHTOKEN_DEBUG" ]; then
   DEBUG_OPT="--debug=$AUTHTOKEN_DEBUG"
fi

bin/authtoken-ws $DBSECURE_OPT $DBHOST_OPT $DBNAME_OPT $DBUSER_OPT $DBPASSWD_OPT $DEBUG_OPT

#
# end of file
#
