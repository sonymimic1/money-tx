#!/bin/sh

set -e

echo "run db migration"
/bin/migrate -path=/migrations/ -database mysql://root:123456@tcp(mysql-test:3306)/transferDB -verbose up

echo "start the app"
exec "$@"