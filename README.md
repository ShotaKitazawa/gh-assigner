# About

<img src="https://github.com/ShotaKitazawa/gh-assigner/blob/images/kairu.png" width="200px">

# Architecture

<img src="https://github.com/ShotaKitazawa/gh-assigner/blob/images/architecture.png" width="400px">

# Migrate DB

* using https://github.com/golang-migrate/migrate
```
migrate -path migrations/ -database 'mysql://root:password@tcp(localhost:3306)/sample' up 4
```
