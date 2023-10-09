<h1>Gin-Gonic Django Authentication Handler</h1>

[![Go](https://github.com/dowling-john/gin_django_auth/actions/workflows/unit_testing.yaml/badge.svg)](https://github.com/dowling-john/gin_django_auth/actions/workflows/unit_testing.yaml)

<h2>If you have a django backend/admin portal this module allows for the connection of you application to the django backend 
authentication and authorisation system. We use the sessionid cookie from the django sessions table to authenticate/authorise
request coming into the gin router.<h2>

<h3>
Features
</h3>
    - Secure Gin routes with a Django session cookie

<h3>Installation</h2>
```sh
    go get github.com/dowling-john/gin_django_auth
```


<h3>Handler Usage</h2>

```golang
    import (
	    "github.com/Flexin1981/gin_django_auth/middleware"
    )

    router.POST("/graphql", middleware.LoginRequired, handlers.GraphQlHandler)
    
```