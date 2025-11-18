# shorten-url

## Develop
### Run the source code
#### MySQL
- The source code requires MySQL. The setup can be see in [docker-compose.yml](./docker/docker-compose.yml)
- To start MySQL. Run command "docker compose -f 'docker/docker-compose.yml' up -d --build 'mysql'" at root folder
- Run SQL script in file [init_database.sql](./docker/init_database.sql) to create database and table

### Application
#### Environment variables
Environment variables used in [.env.sample](./env/.env.sample)
| Variable              | Description                    | Value                       |
| --------------------- | ------------------------------ | --------------------------- |
| SERVICE               | Service name                   | "shorten-url"               |
| HOST                  | Service host                   | "localhost"                 |
| PORT                  | Service port                   | "8080"                      |
| WRITE_TIMEOUT         | Service write response timeout | 15s                         |
| READ_TIMEOUT          | Service read request timeout   | 15s                         |
| PUBLIC_DOMAIN         | Service public domain          | "shortenurl.alwaysdata.net" |
| MIN_SHORTEN_IN_LENGTH | Starting length of encode id   | "6"                         |
| ALLOW_ORIGIN          | Allow origin call to APIs      | "*"                         |
| ALLOW_METHODS         | Allow request methods          | "GET,POST"                  |
| ALLOW_HEADERS         | Allow request headers          | "*"                         |
| MYSQL_HOST            | MySQL host                     | localhost                   |
| MYSQL_PORT            | MySQL port                     | 3306                        |
| MYSQL_USER            | MySQL user                     | root                        |
| MYSQL_PASSWORD        | MySQL password                 | password                    |
| MYSQL_DATABASE        | MySQL database                 | shorten_url                 |

#### Run
- The application is developed with VSCode. The setup for debugging can be seen in [launch.json](./.vscode/launch.json)
- The configuration for the application can be seen in [.env.sample](./env/.env.sample)
- To run the application:
    - Run command "go mod tidy"
    - Create a ".env" file from [.env.sample](./env/.env.sample) in folder env (there is a path to this ".env" file in [launch.json](./.vscode/launch.json))
    - With VSCode:
        - You can use Run and Debug (or Click F5 for ease of use)
    - With terminal:
        - Run command "source ./env/export_env.sh" to export env variables
        - Run "go run ."

### Run tests
go test -v -cover ./internal/adapter/handler ./internal/service

### Postman
[Go to postman guide](./postman/GUIDE.md)

## Build docker image for other platforms
- Start at root of projects
- Execute command "./docker/build_image.sh" to build binary file named "shortlink"
- Execute command "docker image ls | grep "shortlink"" to verify