# go-zap - Seamless integration between Opentracing and Zap

this repo contains a couple possible integrations with go.uber.org/zap


## log

The `log` package allows you to call the same `funcs` as with `zap`, but those will now also log on opentracing if you give a `context` or `span` as parameter.

### Example Usages:

Import the package and use it as your `log` package:
```
import 	"github.com/opentracing-contrib/go-zap/log"
```

Example with Debug level (similar funcs exist for the other levels).
Log on zap only:

```
func Debug(log string, fields ...zapcore.Field)
```

Log on Zap and on Opentracing. The span is given directly or is retrieved from the context:

```
func DebugWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
func DebugWithContext(ctx context.Context, log string, fields ...zapcore.Field)
```


## utils

The `utils` package provides a function to translate a `zap` field into a standard `opentracing` field. Those two are almost similar (in implementation and concept)
