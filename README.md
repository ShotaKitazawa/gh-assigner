# About

![カイル君](https://www.kanatakita.com/images/kairu.png)


# Migrate DB

* using https://github.com/golang-migrate/migrate
```
migrate -path migrations/ -database 'mysql://root:password@tcp(localhost:3306)/sample' up 4
```
