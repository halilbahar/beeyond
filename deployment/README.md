# Deployment

Credit: [caberger/leocloud](https://github.com/caberger/leocloud)

## URLs

| Service | URL |
|-|-|
| Landing Page | <https://student.cloud.htl-leonding.ac.at/beeyond> |
| App (Angular) | <https://student.cloud.htl-leonding.ac.at/beeyond/app/> |
| Backend (Quarkus) | <https://student.cloud.htl-leonding.ac.at/beeyond/api> |
| Validation (Go) | <https://student.cloud.htl-leonding.ac.at/beeyond/valid> |
| Authentication (Keycloak) | <https://student.cloud.htl-leonding.ac.at/beeyond/auth> |

## How to deploy?

Requirements:

* kubectl [kubectl installieren](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/)

Deployment:

* Sign into <https://cloud.htl-leonding.ac.at>
* Install your config file into your ~/.kube folder
* Call `./create-deployment.sh`, to specify following fields pass them as a parameter in the given order. This will merge all .yaml files in the folder [parts](./parts), separated by "---" (IMPORANT: Always have an empty line at the end of the .yaml file).

| Nr | Parameter | Default-Value | Possible Values | Description
|-|-|-|-|-|
| 1 | image-version| latest | * |
| 2 | image-github-acc | halilbahar | * |
| 3 | ingress-path | beeyond | * |

* (optional) Run `kubectl delete -f deployment.yaml` if you already have running deployed it once.
* Run `kubectl apply -f deployment.yaml`.
* With `kubectl get pods` you are able to see if the pods are working as expected.

## Keycloak Configuration - TODO

<https://github.com/keycloak/keycloak-documentation/blob/main/server_admin/topics/admin-cli.adoc>


```bash
kubectl cp ../development-container/keycloak-theme identity-provider-5d5674444b-8nvt4:/opt/jboss/keycloak/themes/beeyond
./kcadm.sh delete realms/school --server http://localhost:8080/auth --realm master --user beeyond
./kcadm.sh create realms -f /tmp/school-realm.json --server http://localhost:8080/auth --realm master --user beeyond
```

### Known Problems

#### X-Frame-Options-Error in Admin Console

```
Refused to display 'https://student.cloud.htl-leonding.ac.at/' in a frame because it set 'X-Frame-Options' to 'deny'.
```

Solution: Because this is a header which is not defined by us, we are not able to remove it.

Workaround: Install an browser extention which ignores the 'X-Frame-Options' header.

| Browser | Link |
|-|-|
| Chrome | [Ignore X-Frame headers](https://chrome.google.com/webstore/detail/ignore-x-frame-headers/gleekbfjekiniecknbkamfmkohkpodhe) |
| Firefox | [Ignore X-Frame-Options Header](https://addons.mozilla.org/de/firefox/addon/ignore-x-frame-options-header/)|

## Useful links

[How to open terminal in K8s-Pod?](https://kubernetes.io/docs/tasks/debug-application-cluster/get-shell-running-container/)

[Kubectl-Spckzettel](https://kubernetes.io/de/docs/reference/kubectl/cheatsheet/)
