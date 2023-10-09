<h1>Gin-Gonic Django Authentication Handler</h1>

<h2>
    Just to serve as a quick note This is currently under development and not ready for production yet.
    Please keep checking back for updates and use the issues to add feature requests
</h2>

If you have a django backend/admin portal this module allows for the connection of you application to the django backend 
authentication and authorisation system. We use the sessionid cookie from the django sessions table to authenticate/authorise
request coming into the gin router.

<h2>Installation</h2>
```bash
    go get github.com/dowling-john/gin_django_auth
```


<h2>Handler Usage</h2>

```golang
    import (
	    "github.com/Flexin1981/gin_django_auth/middleware"
    )

    router.POST("/graphql", middleware.LoginRequired, handlers.GraphQlHandler)
    
```