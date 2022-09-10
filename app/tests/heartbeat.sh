#!/bin/sh
echo "Is api alive?"
echo
API_URL="localhost:10000"
RES=$(curl $SILENT -X GET $API_URL'/api/heartbeat')
ResultError=$(echo $RES | jq -r .err)
echo "Result error:" && echo $ResultError
if [ "$ResultError" != null ]; then printf "${GREEN}PASSED${NC}\n" ; else  printf "${LRED}FAILED${NC}\n" ; exit 1 ;  fi
echo