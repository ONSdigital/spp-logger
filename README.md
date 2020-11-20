# SPP Logger

A collection of logging libraries to implement standardised logging.

## Implementation languages

- [Python](/python)

## Logging spec

### Inter-process communication

- Two compulsory attributes:
    - **log_correlation_id** - This is the correlation value (based on GUID, but may have other details encoded into it)
    - **log_level** - standard python logging levels (Critical/Error/Warning/Info/Debug)
- These attributes are:
    - defined by the initiating function/process.
    - held in a JSON structure
    - held in an attribute called 'context'
    - immutable, and cannot be changed/removed by any subsequent process
- Subsequent processes in the chain can add additional attributes to context as necessary (e.g. reference). These are also immutable.
- All items in the context are printed to the log.

### When logging

- Several compulsory attributes: (NB. these are still very much up for discussion)
    - **timestamp** - Implicit? Legacy question
    - **description** - Bespoke text description of what has happened
    - **user** - Initiating User or "System"
    - **service** - Parent service name. e.g. Validation, rsi-results-pipeline. (Leads to come up with a list)
    - **component** - The lambda/pod/function name. e.g rsi_json_standardisation. 
    - **environment** - Describe where the function is run E.g. dev-sandbox, uat, int, etc. This could include the account-id. (This assumes that log collation doesn't make this implicit)
        - Sandbox
        - Integration
        - UAT
        - Production
    - **deployment** - A specific deployment
        - main
        - bpm-457
        - bpm-587
        - mike
- Logging levels:
    - The standard logging level is '**info**'. (This will capture everything except '**debug**').
        - **Critical**: Always logged. A process has failed. All teams in SPP should be notified. Describes the failure state that just happened.
        - **Error**: Always logged. A process has failed. The owning team in SPP should be notified. Describes the failure state that just happened.
        - **Warning**: Always logged. Something isn't quite right, it hasn't failed, but should be investigated.
        - **Info**: Always logged. Useful information that should always be captured. (e.g. function start/end, or an API call has been made)
        - **Debug**: Sometimes logged. Detailed process information that is useful to trace issues but too voluminous to continually collect.
    - A certain % (e.g. say 5%) of the time the edge/starting function will set the log-level to debug. This percentage should be dynamic if implemented.
        - (This allows us to 'test' the logging and correctly identify what should/shouldn't be handled earlier in the process. It will also often give us enough information to identify issues without having to switch all logging to debug level).
    - Should be determined dynamically unless impossible to do so. There should not be any infrastructure/deployment changes required to update this.
    - Behaviour is treated the same across all environments. i.e. the process for changing the log-level is the same (assuming permitted access, etc)

##  CI

```bash
fly -t xcutting set-pipeline -c ci/pipeline.yml -p spp-logger
```