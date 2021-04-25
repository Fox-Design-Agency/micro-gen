package models

// ReturnGoGitignore will return the byte slice of the generic .gitignore file
func ReturnGoGitignore() ([]byte, error) {
	return []byte(`.DS_Store
*.[56789ao]
*.a[56789o]
*.so
*.pyc
._*
.nfs.*
[56789a].out
*~
*.orig
*.rej
*.exe
.*.swp
core
*.cgo*.go
*.cgo*.c
_cgo_*
_obj
_test
_testmain.go
		
/VERSION.cache
/bin/
/build.out
/doc/articles/wiki/*.bin
/goinstall.log
/last-change
/misc/cgo/life/run.out
/misc/cgo/stdio/run.out
/misc/cgo/testso/main
/src/*.*/
/src/cmd/cgo/zdefaultcc.go
/src/cmd/dist/dist
/src/cmd/go/internal/cfg/zdefaultcc.go
/src/cmd/go/internal/cfg/zosarch.go
/src/cmd/internal/objabi/zbootstrap.go
/src/go/build/zcgo.go
/src/go/doc/headscan
/src/runtime/internal/sys/zversion.go
/src/unicode/maketables
/test.out
/test/garbage/*.out
/test/pass.out
/test/run.out
/test/times.out
		
config
		
# ignore workspace stuffs
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
.history/*
		
/.idea
.project
.classpath
.c9/
*.launch
.settings/
*.sublime-workspace
		
# This file includes artifacts of Go build that should not be checked in.
# For files created by specific development environment (e.g. editor),
# use alternative ways to exclude files from git.
# For example, set up .git/info/exclude or use a global .gitignore.

#temp
secretStuff.go`), nil
}

// ReturnGoREADME will return the byte slice of the generic README file
func ReturnGoREADME() ([]byte, error) {
	return []byte(`# Some Neat Service

This is a pretty neat service that merits a better README
		
## Table of Contents
		
- [Overview](#overview)
- [Getting Started](#getting-started)
	- [Running Locally](#running-locally)
	- [Running In Prod](#running-in-prod)
- [How Do I Add A New Endpoint?](#how-do-i-add-a-new-endpoint)
- [How Do I Add New Functionality?](#how-do-i-add-new-functionality)
- [How Do I Add New Validation?](#how-do-i-add-new-validation)
- [How Do I Add A New SubService?](#how-do-i-add-a-new-subservice)
		
## Overview
		
Should likely put an overview here of the actual service
		
## Getting Started
		
### Running Locally
		
This project requires the following ENV variables:
		
- DBHOST
- DBNAME
- DBUSER
- DBPASS
- PORT: This is the port for the service to run on, the service will default DB port to 5432
- ENVIRONMENT
		
With the ENV variables set, this project can be currently run in isolation by navigating to the run folder and then executing:
		
go run .
		
### Running In Prod
		
Not a thing yet.
		
## How Do I Add A New Endpoint?
		
Adding a new endpoint can be done by going to the main.go file in the ./run folder. Navigate to the correct place on the page noted by comments. Each subservice has its own location for endpoints, please maintain this organization.
		
## How Do I Add New Functionality?
		
New funcationality can be added in a number of way pending the situation. If you need to add a new SubService not currently in the project, please see [How Do I Add A New SubService?](#how-do-i-add-a-new-subservice). Otherwise, navigate to the correct subservice for the new functionality.
		
The order to add the pattern does not matter, as long as all steps are followed. To add new models, simply add the new struct in the respective models file.
		
To add new functions, add the desired function in the controllers. This is the connection point from main.go, the routes, and the other layers of the application. Current controller pattern attempts to follow this pattern:
		
- validate context
- get either query params or parse JSON
- call the db layer through the desired subservice
- call helpers.SendSuccessHeader or helpers.SendErrorHeader
		
Once that is done, procede to the DB layer in the ./pkg/db folder. Add the func definition to the appropiate interface found near the top of the file. Then add the actual func logic below. Try to maintain the order of the func logic with the interface declaration.
		
## How Do I Add New Validation?
		
If more validation is desired, then some things should be considered:
		
- is there currently a validation chain?
- does this require a new validation chain?
		
If there is a current validation chain, then add the desired func to validate a single field. Do not try and validate more than a single field where possible.
		
If a new validation chain is required, check to see if the appropiate func model and validation runner func exists. If not, then create a new one with the desired model pass through, If it does, then simply go to the correct section and add the function that would utilize interface chaining correctly to then run the desired validation/normalization funcs.
		
## How Do I Add A New SubService?
		
If a new subservice is desired, then make the appropiate files, named appropiatly, in the following folders:
		
- ./pkg/route-handlers: Define the new method and the struct for the route-handler
- ./pkg/db: Define the DB interface and DB struct
- ./pkg/validation
- ./pkg/services: Define the new service
		
Once the outline for the service has been set, it then needs to be registered in services.go found in ./pkg/sub-services. Please maintain alphabetical order for the With funcs.
		
Once registered, then go to main.go found in ./run to register the new service and intialize the route-handler if desired. If routes will connect to this new subservice, then make the appropiate place in the routes section and note the section through comments.`), nil
}

// ReturnMigrationSeedFile will return the byte slice of the generic Migrations file
func ReturnGoMigrationSeedFile() ([]byte, error) {
	return []byte(`-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
	
	
	
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back`), nil
}

// ReturnGoDockerignore will return the byte slice of the generic .dockerignore file
func ReturnGoDockerignore() ([]byte, error) {
	return []byte(`/dist
/tmp
/out-tsc
/bazel-out
	
	
/node_modules
	
	
chrome-profiler-events.json
speed-measure-plugin.json
	
	
/.idea
.project
.classpath
.c9/
*.launch
.settings/
*.sublime-workspace
	
	
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
.history/*
	
	
/.sass-cache
/connect.lock
/coverage
/libpeerconnection.log
npm-debug.log
yarn-error.log
testem.log
/typings
	
.DS_Store
Thumbs.db
	
node_modules
.git
.gitignore
	
config`), nil
}

// ReturnGoDockerfile will return the byte slice of the generic Dockerfile file
func ReturnGoDockerfile() ([]byte, error) {
	return []byte(`# Use the official Golang image to create a build artifact.
	# This is based on Debian and sets the GOPATH to /go.
	# https://hub.docker.com/_/golang
	FROM golang:latest as builder
	
	# Create and change to the app directory.
	WORKDIR /app
	
	# Copy local code to the container image.
	COPY . ./
	# Retrieve application dependencies.
	# This allows the container build to reuse cached dependencies.
	COPY go.* ./
	RUN go mod download
	
	WORKDIR /app/run
	# Build the binary.
	RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server
	
	# Use the official Alpine image for a lean production container.
	# https://hub.docker.com/_/alpine
	# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
	FROM alpine:3
	RUN apk add --no-cache ca-certificates
	
	
	# Copy the binary to the production image from the builder stage.
	COPY --from=builder /app/run/server /server
	COPY --from=builder /app/pkg /pkg
	COPY --from=builder /app/resources /resources
	
	EXPOSE 80
	
	# Run the web service on container startup.
	CMD ["/server", "-prod=true"]
	`), nil
}
