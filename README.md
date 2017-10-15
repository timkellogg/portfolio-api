# Portfolio Api

Captures messages from portfolio and saves them in a db.

## Setup

* clone repo
* compile source: `go build`
* execute source: `./portfolio-api`

## Usage

* make curl:

      curl -X POST -d '{ "message": "message", "email": "example@user.com", "name": "John Smith" }' https://localhost:3000/contacts