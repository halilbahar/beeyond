# Dev Images

This folder contains a docker-compose file for development. It starts two containers: KeyCloak and PostgreSQL.

## PostgreSQL

The Postgres container `beeyond-database` provides the **database** `beeyond_db` which has an **user** `beeyond` with the **password** `beeyond`.

## KeyCloak

The KeyCloak container `beeyond-identity-provider` provides the following:

* A realm called `school` with the roles `student` and `teacher`
* A service called `beeyond` with the secret `f203438c-d453-453d-8c66-9e5f22aaf80a`
* 2 users (name-password):
    * `student`-`student` with the role `student`
    * `teacher`-`teacher` with the role `teacher`
