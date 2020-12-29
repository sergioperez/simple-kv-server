# Simple KV Server

## Execute server

```
go run simple-kv-server.go
```

## Env variables:
```
MAX_KEY_AGE_SECONDS: Number of seconds before a key is considered invalid
```


## Set a key:

```
GET /keyname/keyvalue
```

## Get all keys and values:

```
GET /
```
