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


## Note

- use validate lib (ex: https://github.com/go-playground/validator )
