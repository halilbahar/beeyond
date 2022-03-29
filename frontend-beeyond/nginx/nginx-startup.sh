json="{\"backendApiUrl\": \"${BEEYOND_BACKEND_API_URL}\", \"validationApiUrl\": \"${BEEYOND_VALIDATION_API_URL}\", \"keycloakUrl\": \"${BEEYOND_KEYCLOAK_URL}\", \"redirectUrl\": \"${REDIRECT_URL}\"}"
echo $json > /usr/share/nginx/html/app/assets/config.json
