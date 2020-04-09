# Simple go web app

This is a simple demonstration web app written in Go (martini framework) that easily runs on Cloud Foundry.

# Run it

- Just go get

  ```
  go get github.com/yogendra/shorty
  ```

- And run

  ```
  shorty
  ```

- Check it out

  - [http://localhost:8080/twitter][http://localhost:8080/twitter] -> [https://twitter.com/pivotal][https://twitter.com/pivotal]

## Locally

```
go run urls.go main.go
```

## Cloud Foundry

```
cf push
```
