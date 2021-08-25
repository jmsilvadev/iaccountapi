[![Quality And Tests](https://github.com/jmsilvadev/interview-accountapi/actions/workflows/pull-requests.yml/badge.svg)](https://github.com/jmsilvadev/ief/actions/workflows/pull-requests.yml)
[![Release](https://github.com/jmsilvadev/interview-accountapi/actions/workflows/release.yml/badge.svg)](https://github.com/jmsilvadev/ief/actions/workflows/release.yml)

# Candidate

Name: Josué Silva

Email: jmsilvadev@gmail.com

My main language is PHP, but I have been in contact with Go for 2 years, I really like the language and I develop personal projects with it, but I don't use it much in my current company because the team has resistance against the language. For me this is a great opportunity to receive feedback, negative or positive, about my code as I always try to apply the best practices I use in my day-to-day in whatever language I work.

I tried to minimize the use of third-party libs, including tests were done without any third-party libs.

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

# Form3 Take Home Exercise

Engineers at Form3 build highly available distributed systems in a microservices environment. Our take home test is designed to evaluate real world activities that are involved with this role. We recognise that this may not be as mentally challenging and may take longer to implement than some algorithmic tests that are often seen in interview exercises. Our approach however helps ensure that you will be working with a team of engineers with the necessary practical skills for the role (as well as a diverse range of technical wizardry). 

## Instructions
The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with the API. Please note that the fake account API does not require any authorisation or authentication.

A mapping of account attributes can be found in [models.go](./models.go). Can be used as a starting point, usage of the file is not required.

If you encounter any problems running the fake account API we would encourage you to do some debugging first,
before reaching out for help.

## Submission Guidance

### Shoulds

The finished solution **should:**
- Be written in Go.
- Be a client library suitable for use in another software project.
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API.
- Be simple and concise.
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass.

### Should Nots

The finished solution **should not:**
- Use a code generator to write the client library.
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in the test being rejected.
- Use a library for your client (e.g: go-resty). Libraries to support testing or types like UUID are fine.
- Implement client-side validation.
- Implement an authentication scheme.
- Implement support for the fields `data.attributes.private_identification`, `data.attributes.organisation_identification`
  and `data.relationships`, as they are omitted in the provided fake account API implementation.
- Have advanced features, however discussion of anything extra you'd expect a production client to contain would be useful in the documentation.
- Be a command line client or other type of program - the requirement is to write a client library.
- Implement the `List` operation.
> We give no credit for including any of the above in a submitted test, so please only focus on the "Shoulds" above.

## How to submit your exercise

- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, by copying all files you deem necessary for your submission
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License

Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
