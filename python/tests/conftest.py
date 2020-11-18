import logging
from io import StringIO

import pytest

from spp_logger import SPPHandler, SPPLogger, SPPLoggerConfig


@pytest.fixture
def log_stream():
    return StringIO()


@pytest.fixture
def default_handler_config():
    return SPPLoggerConfig(
        service="test-service",
        component="test-component",
        environment="dev",
        deployment="test-deployment",
        user="test-user",
    )


@pytest.fixture
def spp_logger(log_stream, default_handler_config):
    return SPPLogger(
        name="test-logger",
        config=default_handler_config,
        stream=log_stream,
    )


@pytest.fixture
def spp_handler(log_stream, default_handler_config):
    return SPPHandler(
        config=default_handler_config,
        stream=log_stream,
    )


@pytest.fixture
def logger(spp_handler):
    logs = logging.getLogger("test")
    logs.addHandler(spp_handler)
    logs.setLevel(logging.DEBUG)
    return logs


@pytest.fixture
def log_record():
    record = logging.LogRecord(
        name="test",
        level="INFO",
        pathname="pathname",
        lineno=1,
        msg="test",
        args=None,
        exc_info=None,
    )
    record.created = 1605225600
    return record
