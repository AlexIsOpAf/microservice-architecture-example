# Cloud-native Microservice Architecture Example

## What kind of overall service this application provides:

Developing a cloud-native web application taht works on the following problem statement:

- Competitive Call Of Duty gamers are looking for a platform to recruit teammates to form a competitive team for ranked matches
- The Application should support:
    - Creating a team (max 6 players)
    - Managing team-members


## Architecture:

![Architecture](Architecture.png  "Architecture")

- - - -
## Authentication/Authorisation Service:

The application has a user-based access control; You can only gain access to the whole app if you've registered. The service authenticates the user and either returns **200** or **401** for authorised vs unauthorised users respectively.

- - - -
## Database Service: 

The databases will store team details aswell as all player details. UserDB to store all registered users. Following the [Single Database Per Service](https://microservices.io/patterns/data/database-per-service.html).

- - - -

## Team/Player CRUD Service:

This service is self documenting as it's a simple CRUD menu on both Player and Teams.

- - - -

## Project Layout:

```bash
project
├── cmd # Contains microservices
│   ├── '{microservice a}' # Services project which contains "main" functions to start each app
│   │   └── main.go
│   ├── '{microservice b}'
│   │   └── main.go
│   └── '{microservice c}'
│       └── main.go
├── docker #dockerfile for each microservice
│   ├── '{microservice a}'
│   ├── '{microservice b}'
│   └── '{microservice c}'
├── internal
│   ├── '{microservice a}' # internal of each microservice. only visible to its own microservice and cant be shared
│   │   ├── rest # HTTP entry point for REST API
│   │   ├── grpc # gRPC entry
│   │   ├── repo # repository / connector to database
│   │   └── service # business logic layer which is handled by the microservice
│   ├── '{microservice b}'
│   │   ├── rest
│   │   ├── grpc
│   │   ├── repo
│   │   └── service
│   └── '{microservice c}'
│       ├── rest
│       ├── grpc
│       ├── repo
│       └── service
├── pkg # public package directory which can be shared across project / microservices
│   └── models
│      ├── ModelA
│      └── ModelB
├── config
├── README.md
└── .gitignore
```


## Credits:

Credit to [Ayushu](https://www.velotio.com/engineering-blog/build-a-containerized-microservice-in-golang) for laying the groundwork to build up on 

Credit to [kcq & co](https://github.com/golang-standards/project-layout) for some-what standarising go-lang project layout