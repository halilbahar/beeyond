# Dev Images

This folder contains a docker-compose file for development. To create and start the service execute `docker-compose up -d`.
This will produce 2 following images:

* `beeyond/keycloak-dev`
* `beeyond/postgres-dev`

## PostgreSQL (postgres-dev)

This postgres image only has an **user** `beeyond` with the **password** `beeyond` and a **database** `beeyond_db`.

## Keycloak (keycloak-dev)

This keycloak image provides the following:

* A realm called `school` with the roles `student` and `teacher`
* A service called `beeyond` with the secret `f203438c-d453-453d-8c66-9e5f22aaf80a`
* 2 users (name-password):
    * `student`-`student` with the role `student`
    * `teacher`-`teacher` with the role `teacher`
