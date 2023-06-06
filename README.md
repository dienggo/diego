## Golang Base

App Structure : 
```
- app -> all logic of application requirement
    - base -> base app needed
    - controllers -> controller to consume on route
    - dto -> data transfer object (request, response)
    - interfaces -> interface/contract struct needed
    - middlewares -> middleware of the app
    - models -> model of database
    - repositories -> repository to encapsulation query or fetch data
    - services -> to build logic oof the app

- config -> configuration app
- helper -> helper method
- provider -> behavior app needed
- route -> app route (api/web)
```