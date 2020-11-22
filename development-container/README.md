# Dev Images

This folder contains a docker-compose file for development. It starts two containers: KeyCloak and PostgreSQL.

## PostgreSQL

The Postgres container `beeyond-database` provides the **database** `beeyond_db` which has an **user** `beeyond` with the **password** `beeyond`.

## KeyCloak

The KeyCloak container `beeyond-identity-provider` provides the following:

* A realm called `school` with the roles `student` and `teacher`.
* Two services:
  * A service called `beeyond` with the secret `8516ccf3-5ed5-4f42-8718-6a769150ba2a` for the backend application.
  * A service called `beeyond-spa` with PKCE for the frontend application.
* Users (username-password):
    * `lastname`-`firstname` with the role `student`:
      * `bahir`-`halil`
      * `cao`-`sonja`
      * `naderer`-`juliana`
      * `eichhorn`-`moritz`
      * `polleichtner`-`moritz`
      * `sljivic`-`emina`
      * `wallinger`-`marc`
    * `lastname-teacher`-`firstname` with the role `teacher` and the same users as shown above

### How to change the preconfigurations

For creating a realm without any preconfigurations run the following command: 

```shell
docker run --rm -p 8180:8080 --name keycloak -v $(pwd)/tmp:/tmp \
-e KEYCLOAK_USER=admin \
-e KEYCLOAK_PASSWORD=admin \
jboss/keycloak:11.0.2
```

This command creates a directory `/tmp` and starts a KeyCloak container which is available at the address http://localhost:8180. After you have logged in with the username `admin` and the password `admin`, you can create a realm with the name `school`, two services, the users with the roles `student` or `teacher`.

When creating the services you must follow these settings:

* Service `beeyond` for the backend application:
  * The `Access Type` must be `bearer-only`.
* Service `beeyond-spa` for the frontend application:
  * The `Access Type` must be `public`.
  * The `Valid Redirect URIs` must be `http://localhost:4200`.
  * The `Web Origins` must be `+`.
  * Under `Advanced Settings` the `Proof Key for Code` must be `S256` and you can specify the `Access Token Lifespan`.

After you are done with your setup you can run in the `/tmp` directory:

```shell
docker exec -it keycloak /opt/jboss/keycloak/bin/standalone.sh \
-Djboss.socket.binding.port-offset=100 -Dkeycloak.migration.action=export \
-Dkeycloak.migration.provider=singleFile \
-Dkeycloak.migration.realmName=school \
-Dkeycloak.migration.usersExportStrategy=REALM_FILE \
-Dkeycloak.migration.file=/tmp/school-realm.json
```

This will export your configurations as a JSON file. Now you can replace the `school-realm.json` in `/development-container` and change the secret of the `beeyond` service of the JSON file to the one found in `/backend-beeyond/src/main/resources/application.properties`.

After everything is done you can then execute the docker-compose file.
