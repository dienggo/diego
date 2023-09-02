## Diego

App Structure : 
```
- app -> all logic of application requirement
    - controllers -> controller to consume on route
    - dto_request -> data transfer object request
    - dto_response -> data transfer object response
    - middleware -> middleware of the app
    - models -> model of database
    - repositories -> repository to encapsulation query or fetch data
    - services -> to build logic of the app
- config -> configuration app
- pkg -> package usage in app
- provider -> behavior app needed
- route -> app route (api/web)
```

### Application Command Tools
- Run HTTP server
  ```shell
  <your-app> http serve
  ```
  example : `my_app http server`
- Run HTTP server on Development/Local Machine
  ```shell
  go run main.go http serve
  ```

### Installation Diego tools :
Get and install tools from github repo
```shell
go install github.com/dienggo/diego/cmd/diego@v1.3.2
```
Check your GOPATH location, windows -> skip this step
```shell
go env GOPATH
$ <your-location-gopath>
```
Copy path and make alias installed diego, example on macos with **zsh**
```
alias diego="/<your-location-gopath>/bin/diego"
```
Example : 
```
alias diego="/usr/development/go/bin/diego"
```

Use diego
```shell
diego -h
```
Update Diego (**beta** - available at version >= v1.2.0)
```shell
diego update
```

### Diego Command Tools
- #### Build New Project
    ```shell
    diego build awesome_project
    ```
  awesome_project -> your project name & module
- #### Make Controller
    ```shell
    diego generate controller example_case
    ```
    example_case -> your controller & controller stored to `app/controllers/example_case.go`
- #### Make Service
    ```shell
    diego generate service example_case
    ```
    example_case -> your service & service stored to `app/services/example_case.go`
- #### Make Middleware
    ```shell
    diego generate service example_case
    ```
    example_case -> your middleware & middleware stored to `app/middleware/example_case.go`
- #### Make Command
    ```shell
    diego generate cmd example_case
    ```
  example_case -> your command & command stored to `cmd/example_case.go` & also registered on provider/cmd.go automatically. <br/>
  run your command on development/local :
  ```shell
  go run main.go <your-command-name-or-alias>
  ```
  example : `go run main.go my-command`

### Migration Tools
Cause this framework depends on `goose`, install `goose` first ---> https://github.com/pressly/goose <br/>
- Get help usage
  ```shell
  diego migration -h
  ```
- Migrate the DB to the most recent version available
  ```shell
  diego migration up
  ```                   
- Migrate the DB up by 1
  ```shell
  diego migration up-by-one
  ```            
- Migrate the DB to a specific VERSION
  ```shell
  diego migration up-to VERSION
  ```        
- Roll back the version by 1
  ```shell
  diego migration down
  ```                 
- Roll back to a specific VERSION
  ```shell
  diego migration down-to VERSION
  ```      
  example : `migration down-to 123`
- Re-run the latest migration
  ```shell
  diego migration redo
  ```                 
- Roll back all migrations
  ```shell
  diego migration reset
  ```                
- Dump the migration status for the current DB
  ```shell
  diego migration status
  ```               
- Print the current version of the database
  ```shell
  diego migration version
  ```              
- Creates new migration file with the current timestamp
  ```shell
  diego migration create NAME [sql|go]
  ```
  example : `diego migration create your_table sql`
- Apply sequential ordering to migrations
  ```shell
  diego migration fix
  ```                  
- Check migration files without running them
  ```shell
  diego migration validate
  ```             
### Library Usage
```
github.com/go-playground/validator/v10 v10.11.1
github.com/google/uuid v1.3.0
github.com/gorilla/mux v1.8.0
github.com/gorilla/schema v1.2.0
github.com/joho/godotenv v1.4.0
github.com/sirupsen/logrus v1.9.3
github.com/stretchr/testify v1.8.2
github.com/urfave/cli/v2 v2.25.7
golang.org/x/text v0.8.0
gopkg.in/yaml.v2 v2.4.0
gorm.io/driver/mysql v1.3.4
gorm.io/gorm v1.23.7
```
