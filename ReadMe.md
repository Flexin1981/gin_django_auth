<h1>Gin-Gonic Django Authentication Handler</h1>

![Unit Tests](https://github.com/dowling-john/gin_django_auth/actions/workflows/unit_testing.yaml/badge.svg)
![coverage](https://raw.githubusercontent.com/dowling-john/gin_django_auth/badges/.badges/master/coverage.svg)

<h4>
If you have a django backend/admin portal this module allows for the connection of you application to the django backend 
authentication and authorisation system. 
<br><br>
We use the sessionid cookie from the django sessions table to authenticate/authorise
request coming into the gin router.
<h4>

<h3>
Features
</h3>

- Secure Gin routes with a Django session cookie
- Adding login handler to create session on the django server

<h3>Installation</h2>

```
    go get github.com/dowling-john/gin_django_auth
```

<h3>Required environment variables</h3>

```
    GINDJANGOAUTHDBCONNECTIONSTRING : "postgres://<username>:<password>@<host>:5432/<db>?sslmode=disable"
```


<h3>Login Required Usage</h2>

```golang
    import (
	    "github.com/Flexin1981/gin_django_auth/middleware"
    )

    router.POST("/graphql", middleware.LoginRequired, handlers.GraphQlHandler)
    
```

<h3>Handler Usage</h2>

```golang
    import (
	    "github.com/Flexin1981/gin_django_auth/handlers"
    )

    router.POST("/login", handlers.DjangoLoginHandler)
    
```