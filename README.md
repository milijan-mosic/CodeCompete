# CodeCompete

## Run project in development (local) mode

1. Open hosts file in the /etc/ dir.

```sh
sudo nvim /etc/hosts
```

2. Copy and paste these lines of code.

```sh
################################################################

127.0.0.1    app.code-compete.local
127.0.0.1    mongodb.code-compete.local

################################################################
```

3. Run Docker Compose

```sh
docker-compose -f docker-compose.local.yaml up --build --remove-orphans
```

4. Visit the homepage of the app [https://app.code-compete.local](https://app.code-compete.local)
