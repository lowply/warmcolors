#/bin/bash

# for multibyte filenames
export LANG=en_US.UTF-8

LOGDIR="./log"
PUBDIR="./public"
LOGFILE="${LOGDIR}/$(date +%y%m%d_%H%M%S).log"
OPTS="--profile private --no-follow-symlinks --delete"
BUCKET="s3://warmcolors.info"

logger(){
	echo "$(date): [Info] ${1}" | tee -a ${LOGFILE}
}

error(){
	echo "$(date): [Error] ${1}" | tee -a ${LOGFILE} 1>&2 
	exit 1
}

main(){
	type aws > /dev/null 2>&1 || error "aws cli is not installed"
	[ -d ${LOGDIR} ] || mkdir ${LOGDIR}
 	logger "Sync starts"
	aws s3 ${OPTS} sync \
		${PUBDIR} ${BUCKET} | tee -a ${LOGFILE}
 	logger "Sync complete"
}

main "$@"
