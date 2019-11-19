# About

<img src="https://github.com/ShotaKitazawa/gh-assigner/blob/images/kairu.png" width="200px">

# Usage

* prepare config.yml

```
github-user: <github-username>
github-token: <github-token>
slack-bot-id: <slack-bot-id>
slack-bot-token: <slack-bot-token>
slack-channel-ids: <slack-channel-id>,<slack-channel-id>,...
slack-default-channel-id: <slack-channel-id>
```

* run

```
go run main.go --config config.yml
```

# Architecture

<img src="https://github.com/ShotaKitazawa/gh-assigner/blob/images/architecture.png" width="400px">

# Migrate DB

* using https://github.com/golang-migrate/migrate
```
migrate -path migrations/ -database 'mysql://root:password@tcp(localhost:3306)/sample' up 4
```
