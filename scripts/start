#!/bin/sh

##################################
#  This script requires screen!  #
#     sudo apt install screen    #
##################################

echo Executing docker-compose...
cd ../development-container/
docker-compose up -d

echo ----------------------

echo "Starting backend (screen: backend)"
cd ../backend-beeyond/
screen -S backend -X quit
screen -dmS backend ./mvnw clean compile quarkus:dev

# The next step requires the user to execute npm install beforehand!
echo "Starting frontend (screen: frontend)"
cd ../frontend-beeyond/
screen -S frontend -X quit
screen -dmS frontend npm start

echo ----------------------
echo OK!
