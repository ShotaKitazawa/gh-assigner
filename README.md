# About

![カイル君](https://cdn.wikiwiki.jp/to/w/boudai/%E3%82%AB%E3%82%A4%E3%83%AB%E5%90%9B/::attach/kyle-icon_400x400.png)


# Migrate DB

* using https://github.com/golang-migrate/migrate
```
migrate -path migrations/ -database 'mysql://root:password@tcp(localhost:3306)/sample' up 4
```
