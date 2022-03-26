#!/usr/bin/env bash

# create the deployment.yaml file from the parts

HOST=student.cloud.htl-leonding.ac.at
OUTPUT=deployment.yaml
INGRESS_PATH=beeyond
GITHUB_ACCOUNT=halilbahar
VERSION=latest
DEPLOY_KEYCLOAK=false

if [[ $1 ]]; then
    VERSION=$1
fi
if [[ $2 ]]; then
    GITHUB_ACCOUNT=$2
fi
if [[ $3 ]]; then
    INGRESS_PATH=$3
fi
if [[ $3 ]]; then
    INGRESS_PATH=$3
fi
if [[ $4 ]]; then
    DEPLOY_KEYCLOAK=$4
fi

URL=https://$HOST/$INGRESS_PATH
PARTS=$(find ./parts -type f -name "*.yaml" -print | sort)

rm -f $OUTPUT
CNT=0
for file in $PARTS
do
    if [[ $file != *"keycloak"* || $file == *"keycloak"* && $DEPLOY_KEYCLOAK == "true" ]]; then
        if [[ $CNT != "0" ]]
        then
            echo "---" >> $OUTPUT
        fi
        sed -e "s/\$GITHUB_ACCOUNT/$GITHUB_ACCOUNT/g" $file | sed -e "s/\$EMAIL/$INGRESS_PATH/g" | sed -e "s/\$VERSION/$VERSION/g" | sed -e "s;\$HOST;$HOST;g" | sed -e "s;\$URL;$URL;g" >> $OUTPUT
        let CNT+=1
    fi
done

exit 0
