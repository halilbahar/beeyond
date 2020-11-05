# Dev Images

This folder contains a docker-compose file for development. It creates and starts the service which has two containers.

## PostgreSQL

The container `beeyond-database` specifies the PostgreSQL **database** `beeyond` which has an **user** `beeyond` with the **password** `beeyond`.

## KeyCloak

The container `beeyond-identity-provider` specifies the KeyCloak server which has:

* A realm called `school` with the roles `student` and `teacher`
* A service called `beeyond` with the secret `f203438c-d453-453d-8c66-9e5f22aaf80a`
* 2 users (name-password):
    * `student`-`student` with the role `student`
    * `teacher`-`teacher` with the role `teacher`
