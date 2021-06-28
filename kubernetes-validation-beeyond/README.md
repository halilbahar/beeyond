# ![Beeyond](../frontend-beeyond/src/assets/images/beeyond-logo-with-text.png)

**Kubernetes provisioning tool** for the HTBLA Leonding (Dept. of Informatics and Media Technology)

<a href="https://github.com/halilbahar/beeyond/graphs/contributors" alt="Contributors">
  <img src="https://img.shields.io/github/last-commit/halilbahar/beeyond/master"/>
</a>
<a href="https://github.com/halilbahar?tab=packages&repo_name=beeyond" alt="Version">
  <img src="https://img.shields.io/github/v/tag/halilbahar/beeyond"/>
</a>
<a href="https://github.com/halilbahar/beeyond/pulls" alt="PullRequests">
  <img src="https://img.shields.io/github/issues-pr/halilbahar/beeyond"/>
</a>
<a href="https://halilbahar.github.io/beeyond/" alt="Documentation">
  <img src="https://img.shields.io/static/v1?label=docs&message=here&color=orange"/>
</a>
<a href="https://halilbahar.github.io/beeyond/reports.html" alt="Reports">
  <img src="https://img.shields.io/static/v1?label=reports&message=here&color=orange"/>
</a>

## Kubernetes Validation Beeyond

### Getting Started

This is a web API that provides the Kubernetes root definitions and 
validates the YAML that will be deployed 

#### Start DB

Navigate to the [/development-container](../development-container) folder and run 

```shell
docker-compose up -d
```

#### Start project
```shell
go run main.go
```

#### Start tests
```shell
go test ./test/
```


### Endpoint Documentation  
See the swagger docs for futher information of the endpoints [at port 8180/api/swagger-ui](http://localhost:8180/api/swagger-ui).

### How it works
In the frontend a student or teacher can post a YAML file and it gets checked for syntax errors.
Also you can post a constraint on a Kubernetes Object , see 
[Kubernetes JSON definitions](https://kubernetesjsonschema.dev/v1.17.0-standalone-strict/_definitions.json).

# Not complete

