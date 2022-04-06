# gorm-zap
Gorm adapter for zap logger

## Usage

```golang
import "github.com/only1nft/gorm-zap"

db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormzap.New(logger)})
```
