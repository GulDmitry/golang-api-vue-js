Golang API and Vue JS
=====================

## Inside
* [Beego](https://beego.me/)
* [Webpack 2](https://webpack.js.org/)
* [Vue.js](https://vuejs.org) + [vuex](https://vuex.vuejs.org/en/)

## Installation
* `go get github.com/astaxie/beego`
* `go get github.com/beego/bee`
* `go get github.com/satori/go.uuid`
* `go get github.com/smartystreets/goconvey/convey`
* `cp conf/app.conf.dist conf/app.conf`

> **Note**: `echo "export GOPATH=/var/www/golang" >> ~/.zshrc` and `echo "export GOROOT=/usr/lib/go" >> ~/.zshrc`

## Run
* `$GOPATH/bin/bee run`
* Generate docs `$GOPATH/bin/bee generate docs`, run `$GOPATH/bin/bee run -downdoc=true` to download swagger.
* Webpack
  * For `runmode = prod` generate assets `NODE_ENV='production' ./node_modules/.bin/webpack -p`
  * For `runmode = dev` run `./node_modules/.bin/webpack-dev-server --progress --colors --port 8081 --content-base=static/`
* Go to [http://localhost:8080](http://localhost:8080), API [http://localhost:8080/swagger/](http://localhost:8080/swagger/)
* Tests `go test ./...` or `go test ./tests`

## Run Docker
* `docker-compose up` run server and generate API docs
* Webpack
  * For `runmode = prod` generate assets `NODE_ENV='production' ./node_modules/.bin/webpack -p`
  * For `runmode = dev` run `./node_modules/.bin/webpack-dev-server --progress --colors --port 8081 --content-base=static/`
* Go to [http://localhost:8080](http://localhost:8080), API [http://localhost:8080/swagger/](http://localhost:8080/swagger/), MySQL UI [http://localhost:8090](http://localhost:8090)
* Tests `./bin/go test ./...` or `./bin/go test ./tests`

> **Note**: `./bin/go` - make a container and run the command, `./bin/go env` to show containers env variables.

#### TODO
* User.
* Data picker and format data on client.
* ORM.
