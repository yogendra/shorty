# Sorty - Less is More

Simple app to do url shortening. This app is designed to work without any backend. All url mappings are in `urls.go` file. Edit that.

This is a standalne app. Its best to run it on cloudfoundry.

```
git clone https://github.com/yogendra/shorty.git shorty
cd shorty
cf push
```

# Run with docker

```
docker run -p shorty
```

# Running on self-hosted environement with Go

- Just go get

```

go get github.com/yogendra/shorty

```

- And run

```

shorty

```

- Check it out

  - http://localhost:8080/twitter -> https://twitter.com/pivotal

# Build Locally

```

go run urls.go main.go

```
