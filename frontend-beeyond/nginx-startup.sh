json="{\"backendApiUrl\": \"${BEEYOND_BACKEND_API_URL}\", \"validationApiUrl\": \"${BEEYOND_VALIDATION_API_URL}\", \"keycloakUrl\": \"${BEEYOND_KEYCLOAK_URL}\"}"
echo $json > /usr/share/nginx/html/assets/config.json
