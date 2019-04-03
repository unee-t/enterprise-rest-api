#!/bin/bash

STAGE=dev

domain() {
	case $1 in
		prod) echo auroradb.unee-t.com
		;;
		*) echo auroradb.$1.unee-t.com
		;;
	esac
}

ssm() {
if test "$2"
then
	aws --profile $1 ssm get-parameters --names $2 --with-decryption --query Parameters[0].Value --output text
else
	aws --profile ${1:-uneet-dev} ssm describe-parameters
fi
}


show_help() {
cat << EOF
Usage: ${0##*/} [-p]

By default, deploy to dev environment on AWS account 812644853088

	-p          PRODUCTION 192458993663
	-d          DEMO 915001051872

EOF
}

while getopts "pd" opt
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
		*)
			show_help >&2
			exit 1
			;;
	esac
done
AWS_PROFILE=uneet-$STAGE
shift "$((OPTIND-1))"   # Discard the options and sentinel --

echo Connecting to unee_t_enterprise on ${STAGE^^} $(domain $STAGE)

#mysql -s -h $(domain $STAGE) -P 3306 -u $(ssm $AWS_PROFILE UNEE-T_ENTERPRISE_RDS_MASTER_USER) --password=$(ssm $AWS_PROFILE UNEE-T_ENTERPRISE_RDS_MASTER_USER_PASSWORD) unee_t_enterprise
mysql -s -h $(domain $STAGE) -P 3306 -u $(ssm $AWS_PROFILE MYSQL_USER) --password=$(ssm $AWS_PROFILE MYSQL_PASSWORD) unee_t_enterprise
