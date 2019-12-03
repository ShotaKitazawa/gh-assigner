# About

<img src="https://github.com/ShotaKitazawa/gh-assigner/blob/images/kairu.png" width="200px">

# Architecture

<img src="https://github.com/ShotaKitazawa/gh-assigner/blob/images/architecture01.png" width="400px">

<img src="https://github.com/ShotaKitazawa/gh-assigner/blob/images/architecture02.png" width="400px">

# Usage

* prepare config.yml

```
github-user: <github-username>
github-token: <github-token>
slack-bot-id: <slack-bot-id>
slack-bot-token: <slack-bot-token>
slack-channel-ids: <slack-channel-id>,<slack-channel-id>,...
slack-default-channel-id: <slack-channel-id>
google-calendar-id: <google-calendar-id>
crontab: 0 10 * * *
```

* run
    * `config.yaml`: created above
    * `credential.json`: gcp service account key

```
go run main.go --config config.yml --gcp-credential-path credential.json
```

# Migrate DB

* using https://github.com/golang-migrate/migrate
```
migrate -path migrations/ -database 'mysql://root:password@tcp(localhost:3306)/sample' up 4
```
