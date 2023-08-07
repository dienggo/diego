## Diego

App Structure : 
```
- app -> all logic of application requirement
    - controllers -> controller to consume on route
    - dto -> data transfer object (request, response)
    - middlewares -> middleware of the app
    - models -> model of database
    - repositories -> repository to encapsulation query or fetch data
    - services -> to build logic of the app
- config -> configuration app
- pkg -> package usage in app
- provider -> behavior app needed
- route -> app route (api/web)
```

Migration Tools : https://github.com/pressly/goose

### Installation Diego tools :
Get and install tools from github repo
```shell
go install github.com/dienggo/diego/cmd/diego@v1.2.1
```
Check your GOPATH location
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
