Run the server

```
go run ping.go
```

Visit http://0.0.0.0:8080 to see UI.

To adjust warn/error thresholds, change values in `Settings` object in templates/pinger.tmpl.

To test UI, in `ping.go`, use `fakeTime()` (commented out) instead of `ping()`.
