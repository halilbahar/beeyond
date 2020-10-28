# Dev images

This folder contains docker files for development. To build them execute you `./build`. This will produce 2 following images:

* `beeyond/keycloak-dev`
* `beeyond/postgres-dev`

If you want to publish these images or give them other name execute `./build scott`. This will produce again 2 images:

* `scott/keycloak-dev`
* `scott/postgres-dev`

## PostgreSQL (postgres-dev)

This postgres image only has an **user** `beeyond` with the **password** `beeyond` and a **database** `beeyond_db`.

## Keycloak (keycloak-dev)

This keycloak image provides the following:

* A realm called `school` with the roles `student` and `teacher`
* A service called `beeyond` with the secret `f203438c-d453-453d-8c66-9e5f22aaf80a`
* 2 users (name-password):
    * `student`-`student` with the role `student`
    * `teacher`-`teacher` with the role `teacher`
