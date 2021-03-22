# gormotel

OpenTelemetry plugin for [GORM](https://gorm.io)

## Usage

```go
db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{})
if err != nil {
    return err
}
// Use plugin
if err := db.Use(gormotel.Plugin); err != nil {
    return err
}
```