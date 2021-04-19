# MicroService Generator

There are a number of issues with this project and it is not ready for complete use, as there are some manual items that need to be compelted after the template structure is generated.

## Table of Contents

- [Overview](#overview)
- [Things To Do](#things-to-do)

## Overview

This is a very simple microservice genrator that creates the following folder and file structure:

- migrations - placeholder folder to place DB migrations
- pkg - holder folder for most to all of the actual logic
- - db (if db enabled) - holder folder for the database layer for the microservice
- - helpers (if helpers enabled) - holder folder for helper related items, intialized generic success and error http sends
- - middleware - holder folder to place middleware into, currently none are generated
- - models - holder folder for all structs or models
- - route-handlers - holder folder for all route-handlers, which are the contract between the main.go declarations and sub services
- - sub-services - holder folder for initialization logic for sub services within the micro service
- - validation - holder folder for the validation layer that sits on top of the db layer
- run - holder folder for all run related tasks
- - config.go - file for db configuration, currently on Postgres supported
- - main.go - file where the server intialization and routes are declared
- - secretStuff.go - placeholder to inject env vars while in dev
- .dockerignore
- .gitignore
- Dockerfile
- README

## Things To Do

- enable running go mod init and tidy to add the imports by default
- add cli capability so this can be called globally from command line
- change the template write path to be not in run, but this should be done after the thing above
- tons more, but the top three need to be completed before any other sections can
