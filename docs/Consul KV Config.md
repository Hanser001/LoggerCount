## Consul KV Config

summer/user_srv

```json
{
  "name": "user_srv",
  "mysql": {
    "host": "your_host",
    "port": 3306,
    "db": "summer",
    "user": "root",
    "password": "your_password",
    "salt": "nanacho"
  },
  "otel": {
    "endpoint": ":4317"
  },
   "jwt": {
     "key": "your_key"
  }
}
```

summer/file_srv

```json
{
  "name": "file_srv",
  "redis": {
    "host": "your_host",
    "port": 6379,
    "user": "root",
  },
  "otel": {
    "endpoint": ":4317"
  },
  "minio": {
     "endpoint": "127.0.0.1:9000",
     "access_key_id": "your_key",
     "secret_key": "your_key",
     "bucket": "summer"
  },
  "rabbitmq": {
    "host": "127.0.0.1",
    "port": 5672,
    "user": "YOUR_USER",
    "password": "YOUR_PASSWORD",
    "exchange": "summer"
  },
}
```

summer/task_srv

```json
{
  "name": "task_srv",
  "redis": {
    "host": "your_host",
    "port": 6379,
    "user": "root",
  },
  "otel": {
    "endpoint": ":4317"
  },
  "analyze_srv": {
     "name": "analyze_srv"
  }
}
```

summer/analyze_srv

```json
{
  "name": "analyze_srv",
  "otel": {
    "endpoint": ":4317"
  },
  "minio": {
     "endpoint": "127.0.0.1:9000",
     "access_key_id": "your_key",
     "secret_key": "your_key",
     "bucket": "summer"
  },
  "user_srv": {
     "name": "user_srv"
  }
}
```

summer/api_srv

```json
{
    "name": "api",
    "port": 8080,
    "jwt": {
        "key": "your key",
    },
    "otel": {
        "endpoint": ":4317"
    },
    "user_srv": {
        "name": "user_srv"
    },
      "file_srv": {
        "name": "user_srv"
    },
    "task_srv": {
        "name": "user_srv"
    },
    "analyze_srv": {
        "name": "user_srv"
    },
}
```

