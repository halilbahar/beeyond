#!/usr/bin/env bash

# create the deployment.yaml file from the parts

HOST=student.cloud.htl-leonding.ac.at
OUTPUT=deployment.yaml
INGRESS_PATH=beeyond
GITHUB_ACCOUNT=halilbahar
VERSION=latest
KEYCLOAK_REALM=leocloud
KEYCLOAK_CLIENT_ID=beeyond
VOLUMES=false

while [[ $# -gt 0 ]]; do
  case $1 in
    -v|--extension)
      VERSION="$2"
      shift
      shift
      ;;
    -g|--githubaccount)
      GITHUB_ACCOUNT="$2"
      shift
      shift
      ;;
    -i|--ingress)
      INGRESS_PATH="$2"
      shift
      shift
      ;;
    -vo|--volumes)
      VOLUMES=true
      shift
      ;;
    -*|--*)
      echo "Unknown option $1"
      exit 1
      ;;
  esac
done

URL=https://$HOST/$INGRESS_PATH
PARTS=$(find ./parts -type f -name "*.yaml" -print | sort)

rm -f $OUTPUT
CNT=0
for file in $PARTS
do
    if [[ $file != *"volumes"* || $file == *"volumes"* && $VOLUMES == "true" ]]; then
        if [[ $CNT != "0" ]]
        then
            echo "---" >> $OUTPUT
        fi
        sed -e "s/\$GITHUB_ACCOUNT/$GITHUB_ACCOUNT/g" $file | sed -e "s/\$EMAIL/$INGRESS_PATH/g" | sed -e "s/\$VERSION/$VERSION/g" | sed -e "s;\$HOST;$HOST;g" | sed -e "s;\$URL;$URL;g" | sed -e "s;\$KEYCLOAK_CLIENT_ID;$KEYCLOAK_CLIENT_ID;g" | sed -e "s;\$KEYCLOAK_REALM;$KEYCLOAK_REALM;g" >> $OUTPUT
        let CNT+=1
    fi
done

exit 0
