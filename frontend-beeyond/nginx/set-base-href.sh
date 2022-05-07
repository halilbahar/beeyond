# this script changes the <base href="..."> header tag to the environment variable BASE_HREF.
# then the same docker image can be deployed to multiple sub-paths of the same hostname

HTML=/usr/share/nginx/html
if [[ "$BASE_HREF." == "." ]]
then
    echo "no BASE_HREF environment variable set, keep base href... as is is"
else
    cd $HTML
    for file in $(find . -type f -name "*.html"); do
        HREF=$BASE_HREF$(echo $file | cut -d. -f2- | rev | cut -d/ -f2- | rev)
        sed -i "s;<base href=\".*\">;<base href=\"/$HREF/\">;g" $file
        sed -i "s;\$URL_TO_APP;$REDIRECT_URI;g" $file
        sed -i "s;\$URL_TO_DOCUMENTATION;$URL_TO_DOCUMENTATION;g" $file
    done 
fi
