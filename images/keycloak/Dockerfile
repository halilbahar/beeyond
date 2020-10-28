FROM jboss/keycloak:11.0.2

COPY $PWD/school-realm.json /tmp/school-realm.json

ENV KEYCLOAK_USER beeyond
ENV KEYCLOAK_PASSWORD beeyond
ENV DB_VENDOR h2

CMD ["-b", "0.0.0.0", "-Dkeycloak.migration.action=import", "-Dkeycloak.profile.feature.upload_scripts=enabled", "-Dkeycloak.migration.provider=singleFile", "-Dkeycloak.migration.file=/tmp/school-realm.json"]
