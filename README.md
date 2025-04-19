# Distributed Task System
## 2 microservices

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)
![Golang](https://img.shields.io/badge/Golang-v1.23-blue.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Docker](https://img.shields.io/badge/Docker-supported-green.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v17-blue.svg)
![Redis](https://img.shields.io/badge/Redis-v7.2-red.svg)
![Git](https://img.shields.io/badge/Git-v2.39.1-orange.svg)
![Prometheus and Grafana](https://img.shields.io/badge/Prometheus-Grafana-green.svg)
![Help](https://img.shields.io/badge/help-me-brightgreen.svg)
![Project Status](https://img.shields.io/badge/project-active-red.svg)
![Programmer](https://img.shields.io/badge/-I'm%20a%20programmer-yellow.svg)


A small home project for training in microservice architecture and writing applications in Golang.



## Features 

- The whole project and all parts are raised in Docker
- PostgreSQL is used as the database
- Redis is used as a cache
- There are 2 microservices: CRUD for tasks and a service for authorization
- JWT tokens are used in Cookies for security
- All logs and other information are displayed in Grafana, collected using Prometheus
- Microservices communicate through REST
- The project is executed using Clean Architecture, separation of levels, abstractions
- Some parts of the code are covered by unit testing
- Users and Tasks stores in different tables


## Tech

Dillinger uses a number of open source projects to work properly:

- Golang - Backend
- PostgreSQL - DB
- Redis - Cache 
- Prometheus/Grafana - logs, metrics
- Docker - launch
- REST API - A means of communication for microservices
- Clean Code - Architecture Style
- 2 Tables - For task and for users

API Endpoints:

- api-gateway service

  - **GET TASK BY ID:** GET localhost:8080/task/{id}
  - **CREATE TASK:** POST localhost:8080/task
  - **UPDATE TASK BY ID:** PUT localhost:8080/task/{id}
  - **DELETE TASK BY ID:** DELETE localhost:8080/task/{id}
- auth-service service
  
   - **LOGIN:** POST localhost:8081/login
   - **SIGNUP:** POST localhost:8081/signup

You can send the data in the request body. For more information, see the code. A response is returned for each request. If the error is an explanation of why it occurred

## How to Run

1. Clone this repository

```sh
git clone (https://github.com/DmitriyKolesnikM8O/distributed_task_system)
cd distrubuted_task_system
```

2. Assembling images of each of the microservices

```sh
cd api-gateway
docker build -t prod-service:local . 
cd ..
cd auth-service
docker build -t prod-service:local . 
```

3. Launching docker containers

```sh
docker-compose up -d
```

After that, the entire application will be launched

## LICENSE

***MIT*** 

You can use the entire code and the entire project as you see fit. Good Luck! 😊 😊 😊

![The Future is Now](https://img.shields.io/badge/2024-%20The%20Future%20is%20Now-blue.svg)
![Visitor Count](https://komarev.com/ghpvc/?username=DmitriyKolesnikM8O&repo=Distributed-Task-System&label=Visitors&color=007ec6&style=flat-square&abbreviated=true)
