import logging

import immutables
import pytest
from helpers import parse_log_lines

from spp_logger import LogLevelException


def test_logger(spp_logger, log_stream):
    spp_logger.info("my info log message")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 1
    assert log_messages[0]["description"] == "my info log message"


def test_logger_set_context_attribute(spp_logger, log_stream):
    assert spp_logger.context.get("my_attribute") is None
    spp_logger.set_context_attribute("my_attribute", "my_attribute_value")
    assert spp_logger.context.get("my_attribute") == "my_attribute_value"
    spp_logger.info("my info log message")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages[0]) == 11
    assert log_messages[0]["my_attribute"] == "my_attribute_value"


def test_context_can_be_overridden(spp_logger, log_stream):
    spp_logger.set_context(
        immutables.Map(
            log_correlation_id="test",
            log_correlation_type="AUTO",
            log_level=logging.DEBUG,
        )
    )
    spp_logger.info("my first log message")
    spp_logger.set_context(
        immutables.Map(
            log_correlation_id="other test",
            log_correlation_type="AUTO",
            log_level=logging.INFO,
        )
    )
    spp_logger.info("my second log message")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert log_messages[0]["log_correlation_id"] == "test"
    assert log_messages[0]["configured_log_level"] == "DEBUG"
    assert log_messages[1]["log_correlation_id"] == "other test"
    assert log_messages[1]["configured_log_level"] == "INFO"


def test_context_can_be_temporarily_overridden(spp_logger, log_stream):
    spp_logger.set_context(
        immutables.Map(
            log_correlation_id="default_correlation_id",
            log_correlation_type="AUTO",
            log_level=logging.INFO,
        )
    )
    spp_logger.info("my info log message")
    spp_logger.debug("a debug message")
    with spp_logger.override_context(
        immutables.Map(
            log_correlation_id="override_correlation_id",
            log_correlation_type="AUTO",
            log_level=logging.DEBUG,
        )
    ):
        spp_logger.debug("my overridden debug")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 2
    assert log_messages[0]["log_correlation_id"] == "default_correlation_id"
    assert log_messages[0]["description"] == "my info log message"
    assert log_messages[1]["log_correlation_id"] == "override_correlation_id"
    assert log_messages[1]["description"] == "my overridden debug"


def test_setLevel_disabled(spp_logger):
    with pytest.raises(LogLevelException) as err:
        spp_logger.setLevel(logging.WARNING)
    assert str(err.value) == (
        "SPPLogger does not support setting log level this way. "
        + "Please set the log level using the 'log_level' attribute "
        + "on your context"
    )


def test_log_extra_attribute(spp_logger, log_stream):
    spp_logger.info("my info log message", extra={"foobar": "barfoo"})
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 1
    assert log_messages[0]["foobar"] == "barfoo"


def test_log_extra_attribute_cannot_override_context(spp_logger, log_stream):
    spp_logger.info("my info log message", extra={"log_correlation_type": "ERROR"})
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 1
    assert log_messages[0]["log_correlation_type"] == "AUTO"
