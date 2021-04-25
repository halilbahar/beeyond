FROM node AS builder

WORKDIR /app

COPY . .
RUN npm i && npm run build

FROM nginx:alpine

COPY --from=builder /app/dist/beeyond /usr/share/nginx/html/
COPY --from=builder /app/nginx-startup.sh /docker-entrypoint.d/

ENV BEEYOND_BACKEND_API_URL="http://localhost:8080"
ENV BEEYOND_VALIDATION_API_URL="http://localhost:8180/api"
ENV BEEYOND_KEYCLOAK_URL="http://localhost:8280/auth/realms/school"
