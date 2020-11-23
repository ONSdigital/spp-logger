import logging
import sys
from pathlib import Path
from uuid import uuid4

import immutables

MAIN_DIR = Path(__file__).parent.parent
sys.path.insert(0, str(MAIN_DIR))

from spp_logger import SPPLogger, SPPLoggerConfig  # noqa: E402

main_context = immutables.Map(log_correlation_id=str(uuid4()), log_correlation_type="AUTO", log_level=logging.INFO)

secondary_context = immutables.Map(
    log_correlation_id=str(uuid4()), log_correlation_type="AUTO",  log_level=logging.DEBUG
)

config = SPPLoggerConfig(
    service="test-service",
    component="test-component",
    environment="dev",
    deployment="test-deployment",
    user="test-user",
)

# Stream is configurable as any IO, it defaults to stdout
logger = SPPLogger(
    name="my_logger",
    config=config,
    context=main_context,
    stream=sys.stdout,
)

print(f"Starting logger with context: {main_context}\n")
logger.debug("This debug message should not be visible")
logger.info("Got to love an info message")
logger.warning("But be careful, there be dragons!")

with logger.override_context(secondary_context):
    logger.debug("In this context i can print debugs!")

logger.debug("This debug message should not be visible")
