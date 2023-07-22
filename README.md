## Diego

App Structure : 
```
- app -> all logic of application requirement
    - controllers -> controller to consume on route
    - dto -> data transfer object (request, response)
    - interfaces -> interface/contract struct needed
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
go install github.com/daewu14/golang-base/cmd/diego@latest
```
Check your GOPATH location
```shell
go env GOPATH
$ <your-location-gopath>
```
Copy path and make alias installed diego, example on macos with **zsh**
```
alias diego = "/<your-location-gopath>/bin/diego"
```
Use diego
```shell
diego -h
```