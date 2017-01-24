Golang API and Vue JS
=====================

## Inside
* [Beego](https://beego.me/)
* [Webpack 2](https://webpack.js.org/)
* [Vue.js](https://vuejs.org) + [vuex](https://vuex.vuejs.org/en/)

## Installation
* `go get github.com/astaxie/beego`
* `go get github.com/beego/bee`
* `cp conf/app.conf.dist conf/app.conf`

> **Note**: `echo "export GOPATH=/var/www/golang" >> ~/.zshrc` and `echo "export GOROOT=/usr/lib/go" >> ~/.zshrc`

## Run
* `$GOPATH/bin/bee run`
* Generate docs `$GOPATH/bin/bee generate docs`, run `$GOPATH/bin/bee run -downdoc=true` to download swagger.
* Webpack
  * For `runmode = prod` generate assets `NODE_ENV='production' ./node_modules/.bin/webpack -p`
  * For `runmode = dev` run `./node_modules/.bin/webpack-dev-server --progress --colors --port 8081 --content-base=static/`
* Go to [http://localhost:8080](http://localhost:8080)
* API [http://localhost:8080/swagger/](http://localhost:8080/swagger/)

## Tests
* `go test ./...` or `go test ./tests`

#### TODO
* Edit page (vue or static).
* User, session.
* Flash messages, error handling.
* Find out why `$GOPATH/bin/bee run -downdoc=true -gendoc=true` does not work.
* Data picker and format data on client.
