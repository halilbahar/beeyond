json="{\"backendApiUrl\": \"${BEEYOND_BACKEND_API_URL}\", \"validationApiUrl\": \"${BEEYOND_VALIDATION_API_URL}\", \"keycloakUrl\": \"${BEEYOND_KEYCLOAK_URL}\", \"redirectUri\": \"${REDIRECT_URI}\", \"hostUrl\": \"${CLOUD_HOST}\""
echo $json > /usr/share/nginx/html/app/assets/config.json
