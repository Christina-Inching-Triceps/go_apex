# Gopex

apex.tracker.gg API sample

## Develop

### config

Put config.yaml to infrastructure/config/config.yaml  
sample settings is following

```yaml
api:
  apex:
    header: "TRN-Api-Key"
    token: "your-token"
db:
  user: "root"
  password: "root"
  host: "127.0.0.1"
  port: "3306"
  database: "gopex"
```

### TODO/MileStone

1. Repository用の実装 (apiはベタガキ)
1. API用の実装の切り出し
1. テストコードの実装

## Note

- use validate lib (ex: https://github.com/go-playground/validator )
