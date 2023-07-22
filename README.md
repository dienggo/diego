## Golang Base

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