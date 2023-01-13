# bitaksi-matcher
This is a matcher service repository for bitaksi case study.
You can find the `Driver Service` respository link in below

https://github.com/kadirdeniz/bitaksi-driver

[![architecture](https://www.linkpicture.com/q/Screen-Shot-2023-01-13-at-11.19.15.png)](https://www.linkpicture.com/view.php?img=LPic63c11b745529f1064691485)

## Table Of Contents
* [General info](#general-info)
* [Clone the project](#clone-the-project)
* [Setup](#setup)
* [Test](#test)

## General info
This service matches authenticated users with nearest driver using `Driver Service`.In this application JWT based authentication implemented, user should have `authenticated: true` JWT.I tried to follow `Test Driven Development`. `TDD` is a software development approach which test should write before code,then refactoring should be done. I implemented `pactflow` contract test, integration and unit tests and i used `Ginkgo` for clear understanding of test cases. There is a `postman` collection under the `/docs` folder, it can be used for testing.

## Clone the project
```
$ git clone https://github.com/kadirdeniz/bitaksi-matcher.git
$ cd bitaksi-matcher
```

## Setup
##### Application can run with using docker but running locally would be more stable

```
$ make run

$ make dockerized
$ make run-docker
```


 ## Test
 ```
 $ make tests
 ```
