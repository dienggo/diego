## Diego

App Structure : 
```
- app -> all logic of application requirement
    - controllers -> controller to consume on route
    - dto -> data transfer object (request, response)
    - middleware -> middleware of the app
    - models -> model of database
    - repositories -> repository to encapsulation query or fetch data
    - services -> to build logic of the app
- config -> configuration app
- pkg -> package usage in app
- provider -> behavior app needed
- route -> app route (api/web)
```

### Installation Diego tools :
Get and install tools from github repo
```shell
go install github.com/dienggo/diego/cmd/diego@v1.2.6
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

