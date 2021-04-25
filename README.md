# MicroService Generator

This project is likely nearing capability for general use with Golang microservice utilizing Postgres DB.

## Table of Contents

- [Overview](#overview)
- [How to install](#how-to-install)
- [How to Plan](#how-to-plan)
- [Things To Do](#things-to-do)

## Overview

This is a very simple microservice generator that creates the following folder and file structure:

- pkg - holder folder for most to all of the actual logic
  - db (if db enabled) - holder folder for the database layer for the microservice
  - helpers (if helpers enabled) - holder folder for helper related items, initialized generic success and error http sends
  - middleware - holder folder to place middleware into, currently none are generated
  - models - holder folder for all structs or models
  - route-handlers - holder folder for all route-handlers, which are the contract between the main.go declarations and sub services
  - sub-services - holder folder for initialization logic for sub services within the micro service
  - validation - holder folder for the validation layer that sits on top of the db layer
- resources - holder folder for microservice related resources
  - migrations - placeholder folder to place DB migrations
  - postman - placeholder folder to place postman routes
- run - holder folder for all run related tasks
  - config.go - file for db configuration, currently on Postgres supported
  - main.go - file where the server initialization and routes are declared
  - secretStuff.go - placeholder to inject env vars while in dev
- .dockerignore
- .gitignore
- Dockerfile
- README

## How to Install

Right now, the easiest way to install this thing is to clone the repo locally and run a go install from the root.

From the command line, you should be able to then run:

```
micro-gen
```

and the questions will initiate. This process will start to change in the near future. There are currently no flags that are processing.

## How to Plan

The questions are pretty simple:

- Project Name
- Do you need a DB?
  - this will generate the migrations folder
- Do you want helpers?
  - this will generate the generic files that assist with microservice responses
- How Many subservices?
  - this project breaks up the microservice into small sub services that try to focus on a single element, sometimes the subdivision here works out well where the DB dictates how many sub services there will be. Route Handlers can inherit as many sub services as you want, so just plan ahead and divide correctly, else your files will likely become unmanageable.
- sub services questions
  - if you need a db layer and crud are the main considerations to think about for sub services.
  - the db layer will also gen a validation layer that is currently difficult to hook up, more template comments will be added in the future.

## Things To Do

- add placeholder validation funcs
