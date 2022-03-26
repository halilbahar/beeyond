# Deployment

## Keycloak

https://github.com/keycloak/keycloak-documentation/blob/main/server_admin/topics/admin-cli.adoc

```bash
kubectl cp ../development-container/keycloak-theme identity-provider-5d5674444b-8nvt4:/opt/jboss/keycloak/themes/beeyond
./kcadm.sh delete realms/school --server http://localhost:8080/auth --realm master --user beeyond
./kcadm.sh create realms -f /tmp/school-realm.json --server http://localhost:8080/auth --realm master --user beeyond
```
