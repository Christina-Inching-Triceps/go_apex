# Gopex

apex.tracker.gg API sample


## Develop

### Required

- golangci-lint

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
1. go-lintの設定

## Note/Questions

- use validate lib (ex: https://github.com/go-playground/validator )
- ファイル名、クラス名がGo的ではなく長すぎる？
- infrastructureのexternalは外部API用のClientとかを入れるのであって、apiの実装の実態じゃない気がする
  - infrastructureのルート直下にclientを実装してexternal以下の部分でそのclientをインスタンスで持てばいい？
- gorm.db自体はどうやってmockするんだ？？
- returnが*だったりそうじゃなかったりぐちゃぐちゃ

### 命名

- 人物
  - 特定ではない誰か: User
  - データの所持者: Owner
- システム
  - Gopexが主体になるとき: 主語を省略
