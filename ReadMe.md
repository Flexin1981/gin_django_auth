<h1>Gin-Gonic Django Authentication Handler</h1>

<h2>Installation</h2>
```bash
    go get github.com/dowling-john/gin_django_auth
```


<h2>Handler Usage</h2>

```golang
    import (
	    "github.com/Flexin1981/gin_django_auth/middleware")
    )

    router.POST("/graphql", middleware.LoginRequired, handlers.GraphQlHandler)
    
```