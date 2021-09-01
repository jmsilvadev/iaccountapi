[![Quality And Tests](https://github.com/jmsilvadev/interview-accountapi/actions/workflows/pull-requests.yml/badge.svg)](https://github.com/jmsilvadev/ief/actions/workflows/pull-requests.yml)
[![Release](https://github.com/jmsilvadev/interview-accountapi/actions/workflows/release.yml/badge.svg)](https://github.com/jmsilvadev/ief/actions/workflows/release.yml)

## Introduction

The main focus in the package design was to provide a clear solution with a consumer focus and a code with quality and estability, this means that it should provide a clear and simple solution to use. To achieve this, the SOLID standard of single responsibility was followed, all the business logic was encapsulated, having public visibility only of the methods that the consumer needs to use. To grant the quality was use tools of code style and automated tests with 100% of coverage and tests scenarios to prevent known and minimize unkown flaws.

## CI/CD And SemVer

The project uses the Devops concept of continuous integration and continuous delivery through pipelines. To guarantee the CI there is a check inside the pipelines that checks the code quality and runs all the tests. This process generates an artifact that can be viewed and analyzed by developers.

To ensure the continuous delivery, the pipeline uses the automatic semantic versioning creating a versioned realease after each merge in the branch master. This release are tags in the control version system.

The semver system uses the angular commit message model [Angular Commit Message Format](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format).

### NOTE

The docker-compose file has been changed to force the wait of dependent containers until to be healthy and started, so initialization might take a little longer, but it ensures all services are up and running.

```bash
    depends_on:
      postgresql:
        condition: service_healthy
      vault:
        condition: service_started
```

## Helper

To facilitate the developers' work, a Makefile with containerized commands was created.

```bash
➜  interview-accountapi~$ make help
build                          Build docker image
deps                           Install dependencies
doc                            Show package documentation
down                           Stop docker container
fmt                            Applies standard formatting
lint                           Checks Code Style
logs                           Watch docker log files
rebuild                        Rebuild docker container
ssh                            Interactive access to container
test.coverage                  Check project test coverage
test                           Run all available tests
up                             Start docker container in daemon mode
vendor                         Install vendor
vet                            Finds issues in code

```
