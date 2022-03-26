#!/usr/bin/env bash

# create the deployment.yaml file from the parts

HOST=student.cloud.htl-leonding.ac.at
OUTPUT=deployment.yaml
CLOUD_EMAIL=leocloud@htl-leonding.ac.at
GITHUB_ACCOUNT=leocloud
VERSION=latest

if [[ $1 ]]; then
    VERSION=$1
fi
if [[ $2 ]]; then
    GITHUB_ACCOUNT=$2
fi
if [[ $3 ]]; then
    CLOUD_EMAIL=$3
fi


CLOUD_EMAIL=$(echo $CLOUD_EMAIL | sed -e "s/@.*$//")
URL=https://$HOST/$CLOUD_EMAIL
PARTS=$(find ./parts -type f -name "*.yaml" -print | sort)

echo "prepare $OUTPUT for github user $GITHUB_ACCOUNT with base url $CLOUD_EMAIL"

rm -f $OUTPUT
CNT=0
for file in $PARTS
do
    if [[ $CNT != "0" ]]
    then
        echo "---" >> $OUTPUT
    fi
    sed -e "s/\$GITHUB_ACCOUNT/$GITHUB_ACCOUNT/g" $file | sed -e "s/\$EMAIL/$CLOUD_EMAIL/g" | sed -e "s/\$VERSION/$VERSION/g" | sed -e "s;\$HOST;$HOST;g" | sed -e "s;\$URL;$URL;g" >> $OUTPUT
    let CNT+=1
done

cat deployment.yaml
echo "please run now: kubectl apply -f deployment.yaml"
exit 0
