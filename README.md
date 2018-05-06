# Wailingbot
----
Wailingbot is an open source slack bot that implements a Wailing Wall interface.
----

## Slack Configuration

* Create a new app https://api.slack.com/apps
* Name your app and select slash command feature
* Go back to your apps and click on the newly created
  * Save for later the Verification Token
  * Install your app to your workspace
* Configure the slash commands for your app
  ```
  Command       /wwadd
  Request URL   http://mydomain:port/quote

  Command       /wwget
  Request URL   http://mydomain:port/random
  ```

## Bot Configuration
Bot gets confiured thorugh environment variables:
| ENVIRONMENT VARIABLE | DEFAULT VALUE |
| --- | --- |
| PORT | 3000 |
| SLACK_TOKEN |  |
| PG_USER | postgres |
| PG_DBNAME | wailing |
| PG_PASSWORD | postgres |
| PG_HOST | localhost |
| ADD_CMD| /wwadd |
| GET_CMD | /wwget |


## Usage
An usage example is as follows, no other format is supported:
```
/wwadd @slack_user_owning_quote [2017-05-22] This is a legendary quote
/wwget
```

When you get a random quote the response should look like:
```
wailingbot
This is a legendary quote, by @slack_user_owning_quote 2017-05-22
```

The slack command can be named whatever you like:
```
/my_add_command @slack_user_owning_quote [2017-05-22] This is a legendary quote
/my_get_cmmmand
```
To do so you have to configure it properly and make sure it points to the right endpoints in the command section:
```
Command       /my_add_command
Request URL   http://mydomain:port/quote

Command       /my_get_command
Request URL   http://mydomain:port/random
```

## To build golang Wailingbot artifact

If you want to build Wailingbot right away there:
```
$ docker-compose build
```

## To test Wailingbot

First of all build the docker images within the docker compose for testing:
```
$ docker-compose build
```

Run unit tests suite:
```
$ docker-compose run unit_tests go test ./... -tags=unit
$ docker-compose down
```

Run integration tests suite (this depends on the database and the wailingbot containers):
```
$ docker-compose run integration_tests go test ./... -tags=integration
$ docker-compose down
```

## To run Wailingbot in a docker-compose
```
$ export WW_SLACK_TOKEN=abcdefghijk
$ docker-compose up -d
```
