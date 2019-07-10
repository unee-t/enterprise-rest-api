#!/bin/bash

STAGE=dev

ssm() {
	aws --profile $1 ssm get-parameters --names $2 --with-decryption --query Parameters[0].Value --output text
}

show_help() {
cat << EOF
Usage: ${0##*/} [-p]

By default, deploy to dev environment on AWS account 812644853088

	-p          PRODUCTION 192458993663
	-d          DEMO 915001051872
	-l          localhost

EOF
}

while getopts "pdl" opt
do
	case $opt in
		p)
			echo "PRODUCTION" >&2
			STAGE=prod
			;;
		d)
			echo "DEMO" >&2
			STAGE=demo
			;;
		l)
			echo "localhost" >&2
			STAGE=localhost
			;;
		*)
			show_help >&2
			exit 1
			;;
	esac
done
AWS_PROFILE=uneet-$STAGE
shift "$((OPTIND-1))"   # Discard the options and sentinel --

export MYSQL_ROOT_PASSWORD=$(ssm $AWS_PROFILE MYSQL_ROOT_PASSWORD)
export MYSQL_HOST=$(ssm $AWS_PROFILE MYSQL_HOST)

output=./sql/prime.sql
#output=$STAGE-enterprise-backup-$(date +%s).sql

# Get structure, no data
mysqldump --no-data --set-gtid-purged=OFF --single-transaction --skip-lock-tables --column-statistics=0 -R \
	-h $MYSQL_HOST -P 3306 -u root --password=$MYSQL_ROOT_PASSWORD unee_t_enterprise > $output
# Only get data from tables we need
mysqldump --set-gtid-purged=OFF --column-statistics=0 -h $MYSQL_HOST -P 3306 -u root --password=$MYSQL_ROOT_PASSWORD \
	unee_t_enterprise ut_unit_types property_groups_countries ut_map_external_source_users unte_api_keys ut_map_external_source_units >> $output
