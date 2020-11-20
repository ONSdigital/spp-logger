# SPP Logger - Python

A python log handler to ship standardised logs.

## Installation

`pip install git+https://github.com/ONSdigital/spp-logger.git`

## Usage

```python
from spp_logger import SPPLogger, SPPLoggerConfig

config = SPPLoggerConfig(
    service="test-service",
    component="test-component",
    environment="dev",
    deployment="test-deployment",
)
logger = SPPLogger(
    name="test-logger",
    config=config,
)
```


### SPPLoggerConfig

**Args**:

| Argument    | Type | ENV var           | Default                                                                      |
|-------------|------|-------------------|------------------------------------------------------------------------------|
| service     | str  | `SPP_SERVICE`     | N/A                                                                          |
| component   | str  | `SPP_COMPONENT`   | N/A                                                                          |
| environment | str  | `SPP_ENVIRONMENT` | N/A                                                                          |
| deployment  | str  | `SPP_DEPLOYMENT`  | N/A                                                                          |
| user        | str  | `SPP_USER`        | `None` - When set to None the handler will detect the current logged in user |
| timezone    | str  | `TIMEZONE`        | `UTC`                                                                        |

#### Load from env

```python
from spp_logger import SPPLoggerConfig

config = SPPLoggerConfig.from_env()
```

### SPPLogger

| Argument  | Type            | Default                                                                                                                               |
|-----------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------|
| name      | str             | N/A                                                                                                                                   |
| config    | SPPLoggerConfig | `UTC`                                                                                                                                 |
| context   | immutables.Map  | `None` - Will auto generate a context in the form `{"log_correlation_id": uuid, "log_correlation_type": type, log_level": log_level}` |
| log_level | int             | `logging.INFO` - This is only used if `context` is set to `None`                                                                      |
| stream    | IO              | `sys.stdout`                                                                                                                          |

### SPPHandler

| Argument  | Type            | Default                                                                                                                               |
|-----------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------|
| config    | SPPLoggerConfig | `UTC`                                                                                                                                 |
| context   | immutables.Map  | `None` - Will auto generate a context in the form `{"log_correlation_id": uuid, "log_correlation_type": type, log_level": log_level}` |
| log_level | int             | `logging.INFO` - This is only used if `context` is set to `None`                                                                      |
| stream    | IO              | `sys.stdout`                                                                                                                          |

### Context

A context must be an immutable Map with the properties `log_correlation_id`, `log_correlation_type` and `log_level`.

The intention of a context is that the initialising process will configure it and pass it down to any other
initialisations. As a result of this logs can be correlated using the `log_correlation_id` and the `log_level`
is auto set based on the top level initialisation. This works particularly well with serverless constructs where
your module and function is effectively the same thing.

Example:

```python
from spp_logger import SPPLogger, SPPLoggerConfig

config = SPPLoggerConfig.from_env()
logger = SPPLogger(
    name="test-logger",
    config=config,
)

def handler(event, context):
    logger.info("Lambda started")
    return "I ran my lambda"
```

Result:
```json
{"log_level":"INFO","timestamp":"2020-11-17T12:14:58.990943+00:00","description":"Lambda started","user":"test-user","service":"test-service","component":"test-component","environment":"dev","deployment":"test-deployment","log_correlation_id":"e00b4eb1-a853-4955-b38a-fb4a5ea305e4","configured_log_level":"INFO"}
```

However this may not be the desired behaviour in long running app as your `context` is separate
from your apps lifecycle.

In this instance it is recommended that you pass your desired context to your calling application.

Example:

```python
def my_function(context, my_var):
    with logger.override_context(context):
        logger.info("Started my_function")
    return my_var
```

#### Updating context attributes

Context attributes are immutable, however it is possible to add new ones. All attributes on a context will appear in the logs.

```python
logger.set_context_attribute("my_new_attribute", "my_attribute_value")
```
