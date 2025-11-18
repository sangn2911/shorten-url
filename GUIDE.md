# shorten-url

## Develop

### Run the Source Code

#### MySQL
- The source code requires MySQL. The setup can be found in [docker-compose.yml](./docker/docker-compose.yml).  
- To start MySQL, run the following command from the project root:  
    ```bash
    docker compose -f "docker/docker-compose.yml" up -d --build mysql
    ```  
- Run the SQL script in [init_database.sql](./docker/init_database.sql) to create the database and tables.  

---

### Application

#### Environment Variables
Environment variables are defined in [.env.sample](./env/.env.sample):

| Variable              | Description                    | Value                       |
| --------------------- | ------------------------------ | --------------------------- |
| SERVICE               | Service name                   | "shorten-url"               |
| HOST                  | Service host                   | "localhost"                 |
| PORT                  | Service port                   | "8080"                      |
| WRITE_TIMEOUT         | Service write response timeout | 15s                         |
| READ_TIMEOUT          | Service read request timeout   | 15s                         |
| PUBLIC_DOMAIN         | Service public domain          | "shortenurl.alwaysdata.net" |
| MIN_SHORTEN_IN_LENGTH | Starting length of encode ID   | "6"                         |
| ALLOW_ORIGIN          | Allowed origins for API calls  | "*"                         |
| ALLOW_METHODS         | Allowed request methods        | "GET,POST"                  |
| ALLOW_HEADERS         | Allowed request headers        | "*"                         |
| MYSQL_HOST            | MySQL host                     | localhost                   |
| MYSQL_PORT            | MySQL port                     | 3306                        |
| MYSQL_USER            | MySQL user                     | root                        |
| MYSQL_PASSWORD        | MySQL password                 | password                    |
| MYSQL_DATABASE        | MySQL database                 | shorten_url                 |

#### Run
- The application is developed using VSCode. Debugging setup is in [launch.json](./.vscode/launch.json).  
- Configuration is managed via [.env.sample](./env/.env.sample).  

**Steps to run the application:**

1. Install dependencies:
    ```bash
    go mod tidy
    ```
2. Create a `.env` file from [.env.sample](./env/.env.sample) in the `env` folder (the path is referenced in [launch.json](./.vscode/launch.json)).  

**With VSCode:**
- Use **Run and Debug** or press **F5** for convenience.  

**With Terminal:**
- Export environment variables:
    ```bash
    source ./env/export_env.sh
    ```
- Run the application:
    ```bash
    go run .
    ```

#### Run Tests
```bash
go test -v -cover ./internal/adapter/handler ./internal/service
```

---

### Postman
- [Go to Postman guide](./postman/GUIDE.md)

---

## Deploy on alwaysdata
- [Go to deployment guide](./alwaysdata/DEPLOYMENT.md)

---

## Build Docker Image for Other Platforms
- Start at the project root.  
- Build the binary `shortlink` using:  
    ```bash
    ./docker/build_image.sh
    ```
- Verify the Docker image:  
    ```bash
    docker image ls | grep "shortlink"
    ```
