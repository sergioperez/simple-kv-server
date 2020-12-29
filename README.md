# Simple KV Server

## Execute server

```
go run simple-kv-server.go
```

## Env variables:
```
MAX_KEY_AGE_SECONDS: Number of seconds before a key is considered invalid
ENABLE_EXPORT_METRICS: Bool - Enables the /metrics endpoint to show how many values are in each key
```


## Set a key:

```
GET /keyname/keyvalue
```

## Get all keys and values:

```
GET /
```
