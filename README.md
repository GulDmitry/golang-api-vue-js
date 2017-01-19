Golang API and Vue JS frontend
==============================

## Installation
* `go get github.com/astaxie/beego`
* `go get github.com/beego/bee`
* `cp conf/app.conf.dist conf/app.conf`

> **Note**: `echo "export GOPATH=/var/www/golang" >> ~/.zshrc` and `echo "export GOROOT=/usr/lib/go" >> ~/.zshrc`

## Run
* `$GOPATH/bin/bee run`
* Generate docs `$GOPATH/bin/bee generate docs`, run `$GOPATH/bin/bee run -downdoc=true` to download swagger.
* Go to [http://localhost:8080](http://localhost:8080)
* Docs [http://localhost:8080/swagger/](http://localhost:8080/swagger/)

#### TODO
* Webpack + vue.
* Sortable table + CRUD.
* Model and controller tests.
* Find out why `$GOPATH/bin/bee run -downdoc=true -gendoc=true` cannot generate docs.
