set -o allexport
source .env set
# +o allexport

$HOME/cloud-sql-proxy $GOOGLE_INSTANCE_CONNECTION_NAME -p $DB_PORT --credentials-file=$GOOGLE_APPLICATION_CREDENTIALS &