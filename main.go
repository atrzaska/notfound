package main

import (
	"github.com/gin-gonic/gin"
  "net/http"
)

var body string = `
<!DOCTYPE html>
<html>
    <head>
        <title>ERROR 404 - Not Found</title>

        <style type="text/css">
            html, body {
                height: 100%;
                background-color: #666;
                font-family: Helvetica, sans-serif
            }

            body {
                color: #fff;
                text-align: center;
                text-shadow: 0 1px 3px rgba(0,0,0,.5);
            }

            h1 {
                font-size: 58px;
                margin-top: 20px;
                margin-bottom: 10px;
                font-family: inherit;
                font-weight: 500;
                line-height: 1.1;
                color: inherit;
            }

            .site-wrapper {
                display: table;
                width: 100%;
                height: 100%;
                min-height: 100%;
            }

            .site-wrapper-inner {
                display: table-cell;
                vertical-align: top;
            }

            .cover-container {
                margin-right: auto;
                margin-left: auto;
            }

            .site-wrapper-inner {
                vertical-align: middle;
            }
            .cover-container {
                width: 100%;
            }
        </style>
    </head>
    <body>
        <div class="site-wrapper">
          <div class="site-wrapper-inner">
            <div class="cover-container">
                <h1 class="cover-heading">ERROR 404</h1>
                <h5>PAGE NOT FOUND</h5>
            </div>
          </div>
        </div>
    </body>
</html>
`

func main() {
  gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/*all", func(c *gin.Context) {
    c.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte(body))
  })
  router.Run(":4400")
}
