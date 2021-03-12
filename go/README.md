# SPP Logger - Go

A go log handler to ship standardised logs.

## Installation

`go get github.com/ONSDigital/spp-logger/go/spp_logger`

## Usage

### Quickstart

When creating an instance of the logger it will need a `config`, a `context`, a `log_level` and a `stream` 

```go
import (
	"os"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
)

func main() {
    context, _ := spp_logger.NewContext("INFO", "correlation id")
    
    config := spp_logger.Config{
            Service:     "test_service",
            Component:   "test_component",
            Environment: "test_environment",
            Deployment:  "test_deployment",
            Timezone:    "UTC",
        }
        
    logger, _ := spp_logger.NewLogger(config, context, "WARNING", os.Stdout)
    
    logger.info("Function started")
}
```


### LoggerConfig

**Args**:

| Argument    | Type | ENV var           | Default                                                                      |
|-------------|------|-------------------|------------------------------------------------------------------------------|
| service     | str  | `SPP_SERVICE`     | N/A                                                                          |
| component   | str  | `SPP_COMPONENT`   | N/A                                                                          |
| environment | str  | `SPP_ENVIRONMENT` | N/A                                                                          |
| deployment  | str  | `SPP_DEPLOYMENT`  | N/A                                                                          |
| timezone    | str  | `TIMEZONE`        | `UTC`                                                                        |


### Logger

| Argument  | Type              | Default                                                                                                                               |
|-----------|-------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| config    | SPPLoggerConfig   | `UTC`                                                                                                                                 |
| context   | map[string]string | `None` - Will generate a context in the form `{"log_correlation_id": uuid, log_level": log_level}` |
| log_level | int               | `logging.INFO` - This is only used if `context` values are empty                                                      |
| stream    | IO                | `sys.stdout`                                                                                                                          |

The logger contains a `go_log_level`, this is because there is no critical level in go, so we set a hook which is actually an error log. This means when a critical log is called the `go_log_level` will be `ERROR` wheras the `log_level` will be critical. This is to be in keeping with the other loggers.

**Note**: If an attribute name overlaps with a context, the context always takes preference.
                                                                                                        |

### Context

A context must be a map[string]string and must contain the properties `log_correlation_id` and `log_level`.

The intention of a context is that the initialising process will configure it and pass it down to any other
initialisations. As a result of this logs can be correlated using the `log_correlation_id` and the `log_level`
is auto set based on the top level initialisation. This works particularly well with serverless constructs where
your module and function is effectively the same thing.

#### Example set context with additional fields, these additional fields will be logged:

```go
context := map[string]string{"logLevel": "INFO", "correlationID": "test_id", "survey": "survey", "period": "period"}

```
#### Example new context with log_correlation_id and log_level passed in:
```go
context, _ := spp_logger.NewContext("DEBUG", "uuid.NewString()")

```
#### Example new context with no parameters passed in, returns and INFO level and a random uuid log_correlation_id:

```go
context, _ := spp_logger.NewContext("", "")

```

#### Updating context attributes

Context attributes are immutable, however it is possible to add new ones. All attributes on a context will appear in the logs.

```go
logger.SetContextAttribute("my_new_attribute", "my_attribute_value")
```

#### Returning context
```go
logger.Context()
```

Log Result:
```json
{"component":"test_component","configured_log_level":"INFO","log_correlation_id":"correlation id","deployment":"test_deployment","description":"Got to love an info message","environment":"test_environment","go_log_level":"info","log_level":"INFO","service":"test_service","timestamp":"2021-02-22T10:46:17+00:00","timezone":"UTC"}
```

However this may not be the desired behaviour in long running app as your `context` is separate
from your apps lifecycle.

In this instance it is recommended that you pass your desired context to your calling application.

Example:

```go
func my_function(context){
    logger.override_context(context)
    logger.info("Started my_function")
    ...
}
```

