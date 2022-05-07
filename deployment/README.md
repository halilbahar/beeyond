# Deployment

Credit: [caberger/leocloud](https://github.com/caberger/leocloud)

## URLs

| Service | URL |
|-|-|
| Landing Page | <https://student.cloud.htl-leonding.ac.at/beeyond> |
| App (Angular) | <https://student.cloud.htl-leonding.ac.at/beeyond/app/> |
| Backend (Quarkus) | <https://student.cloud.htl-leonding.ac.at/beeyond/api> |
| Validation (Go) | <https://student.cloud.htl-leonding.ac.at/beeyond/valid> |
| Authentication (Keycloak) | <https://auth.cloud.htl-leonding.ac.at/auth/> |

## How to deploy?

Requirements:

* kubectl [kubectl installieren](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/)

Deployment:

* Sign into <https://cloud.htl-leonding.ac.at>
* Install your config file into your ~/.kube folder
* Call `./create-deployment.sh`, to specify following fields pass them as a parameter in the given order. This will merge all .yaml files in the folder [parts](./parts), separated by "---" (IMPORANT: Always have an empty line at the end of the .yaml file).

| Parameter | Default-Value | Possible Values | Description
|-|-|-|-|
| -v \<value>, --version \<value> | latest | * | Sets the image version (same for all images) |
| -g \<value>, --githubaccount \<value>  | halilbahar | * | Sets the owner of the images |
| -i \<value>, --ingress \<value> | beeyond | * | Sets the ingress path |
| -vo, --volumes | false | no values | Sets if the volumes should be includes as well |

* (optional) Run `kubectl delete -f deployment.yaml` if you already have running deployed it once.
* Run `kubectl apply -f deployment.yaml`.
* With `kubectl get pods` you are able to see if the pods are working as expected.

## Keycloak

### Known Problems

#### X-Frame-Options-Error in Admin Console

```text
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
