# MicroService Generator

There are a number of issues with this project and it is not ready for complete use, as there are some manual items that need to be compelted after the template structure is generated.

## Table of Contents

- [Overview](#overview)
- [How to Plan](#how-to-plan)
- [Things To Do](#things-to-do)

## Overview

This is a very simple microservice genrator that creates the following folder and file structure:

- migrations - placeholder folder to place DB migrations
- pkg - holder folder for most to all of the actual logic
  - db (if db enabled) - holder folder for the database layer for the microservice
  - helpers (if helpers enabled) - holder folder for helper related items, intialized generic success and error http sends
  - middleware - holder folder to place middleware into, currently none are generated
  - models - holder folder for all structs or models
  - route-handlers - holder folder for all route-handlers, which are the contract between the main.go declarations and sub services
  - sub-services - holder folder for initialization logic for sub services within the micro service
  - validation - holder folder for the validation layer that sits on top of the db layer
- run - holder folder for all run related tasks
  - config.go - file for db configuration, currently on Postgres supported
  - main.go - file where the server intialization and routes are declared
  - secretStuff.go - placeholder to inject env vars while in dev
- .dockerignore
- .gitignore
- Dockerfile
- README

This template will be created in the run folder currently, but this will change soon.
Once the project is created, move it to where you would like, then navigate in the
terminal to the project and go to the root, then go ahead and run:

```
go mod init <micro service name>
```

After that, run:

```
go mod init <micro service name>
```

Once those have been run, you will need to manualy:

- update the imports in the files
- add the models to make validation happy
- add in actual info into the secretStuff.go file

You should be good to go then with the basic pattern, just navigate to the /run folder of the project and run:

```
go run .
```

## How to Plan

The questions are pretty simple:

- Project Name
- Do you need a DB?
  - this will generate the migrations folder
- Do you want helpers?
  - this will generate the generic files that assist with microservice responses
- How Many subservices?
  - this project breaks up the microservice into small sub services that try to focus on a single element, sometimes the subdivision here works out well where the DB dictates how many sub services there will be. Route Handlers can inherit as many sub services as you want, so just plan ahead and divide correcly, else your files will likely become unmanagable.
- sub services questions
  - if you need a db layer and crud are the main considerations to think about for sub services.
  - the db layer will also gen a validation layer that is currently difficult to hook up, more template comments will be added in the future.

## Things To Do

- enable running go mod init and tidy to add the imports by default
- add cli capability so this can be called globally from command line
- change the template write path to be not in run, but this should be done after the thing above
- tons more, but the top three need to be completed before any other sections can
