FROM node AS builder

WORKDIR app
RUN pwd

COPY ./app .
RUN npm i && npm run build --base="/"

FROM nginx:alpine

COPY landing-page /usr/share/nginx/html
COPY --from=builder /app/dist/beeyond /usr/share/nginx/html/app
COPY nginx/nginx-startup.sh /docker-entrypoint.d/
COPY nginx/set-base-href.sh /docker-entrypoint.d/
COPY nginx/default.conf /etc/nginx/conf.d/default.conf

ENV BEEYOND_BACKEND_API_URL="http://localhost:8080"
ENV BEEYOND_VALIDATION_API_URL="http://localhost:8180/api"
ENV BEEYOND_KEYCLOAK_URL="http://localhost:8280/auth/realms/school"
