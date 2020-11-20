import logging
from unittest import mock
from uuid import uuid4

import immutables
import pytest
from freezegun import freeze_time
from helpers import is_json, is_valid_uuid, parse_log_lines

from spp_logger import ContextError, ImmutableContextError, SPPHandler


@freeze_time("2020-11-13")
def test_handler_logs(logger, log_stream):
    logger.info("my info log message")
    log_line = log_stream.getvalue()

    assert (
        is_json(log_line) is True
    ), f"Expected log lines to be JSON but was: '{log_line}'"
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 1
    assert len(log_messages[0]) == 10
    assert log_messages[0]["log_level"] == "INFO"
    assert log_messages[0]["timestamp"] == "2020-11-13T00:00:00+00:00"
    assert log_messages[0]["description"] == "my info log message"
    assert log_messages[0]["service"] == "test-service"
    assert log_messages[0]["component"] == "test-component"
    assert log_messages[0]["environment"] == "dev"
    assert log_messages[0]["deployment"] == "test-deployment"
    assert log_messages[0]["user"] == "test-user"
    assert log_messages[0]["configured_log_level"] == "INFO"
    assert is_valid_uuid(log_messages[0]["log_correlation_id"])


def test_handler_multiline_logs(logger, log_stream):
    logger.info("my info log message\nwith an extra line")
    logger.info("a second log message")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert log_messages[0]["description"] == "my info log message\nwith an extra line"
    assert log_messages[1]["description"] == "a second log message"


def test_get_timestamp(spp_handler, log_record):
    assert spp_handler.get_timestamp(log_record) == "2020-11-13T00:00:00+00:00"


def test_get_user(spp_handler):
    assert spp_handler.get_user() == "test-user"


@mock.patch("getpass.getuser")
def test_get_user_dynamic(mock_get_user, spp_handler):
    mock_get_user.return_value = "my_test_user"
    spp_handler.config.user = None
    assert spp_handler.get_user() == "my_test_user"
    spp_handler.get_user()
    spp_handler.get_user()
    mock_get_user.assert_called_once()


def test_context_is_immutable(default_handler_config, log_stream):
    log_handler = SPPHandler(
        config=default_handler_config,
        context=immutables.Map(
            log_correlation_id=str(uuid4()), log_level=logging.WARNING
        ),
        stream=log_stream,
    )
    assert log_handler.context["log_level"] == logging.WARNING
    with pytest.raises(Exception) as err:
        log_handler.context["log_level"] = "foobar"
    assert (
        str(err.value)
        == "'immutables._map.Map' object does not support item assignment"
    )


def test_context_can_be_overridden(logger, spp_handler, log_stream):
    spp_handler.set_context(
        immutables.Map(log_correlation_id="test", log_level=logging.DEBUG)
    )
    logger.info("my first log message")
    spp_handler.set_context(
        immutables.Map(log_correlation_id="other test", log_level=logging.INFO)
    )
    logger.info("my second log message")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert log_messages[0]["log_correlation_id"] == "test"
    assert log_messages[0]["configured_log_level"] == "DEBUG"
    assert log_messages[1]["log_correlation_id"] == "other test"
    assert log_messages[1]["configured_log_level"] == "INFO"


def test_set_context_attribute(logger, spp_handler, log_stream):
    spp_handler.set_context_attribute("my_attribute", "my_attribute_value")
    assert spp_handler.context.get("my_attribute") == "my_attribute_value"
    logger.info("my info log message")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages[0]) == 11
    assert log_messages[0]["my_attribute"] == "my_attribute_value"


def test_set_context_attribute_update(spp_handler):
    with pytest.raises(ImmutableContextError) as err:
        spp_handler.set_context_attribute("log_level", "DEBUG")
    assert (
        str(err.value)
        == "Context attributes are immutable, could not override 'log_level'"
    )


def test_log_level_set_by_context(spp_handler, log_stream):
    spp_handler.set_context(
        immutables.Map(log_correlation_id="TEST", log_level=logging.ERROR)
    )
    logger = logging.getLogger("test_log_level_set_by_context")
    logger.addHandler(spp_handler)
    logger.info("my info log message")
    logger.debug("my debug log message")
    logger.warning("this is a warning")
    logger.critical("this is a critical warning")
    logger.error("this is an error")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 2
    assert log_messages[0]["log_level"] == "CRITICAL"
    assert log_messages[0]["configured_log_level"] == "ERROR"
    assert log_messages[1]["log_level"] == "ERROR"
    assert log_messages[1]["configured_log_level"] == "ERROR"


def test_context_must_be_immutable(default_handler_config):
    with pytest.raises(ImmutableContextError) as err:
        SPPHandler(
            config=default_handler_config,
            context=dict(log_correlation_id=str(uuid4()), log_level=logging.WARNING),
        )
    assert str(err.value) == "Context must be a type of 'immutables.Map'"


def test_context_must_be_immutable_when_overridden(spp_handler):
    with pytest.raises(ImmutableContextError) as err:
        spp_handler.set_context(
            dict(log_correlation_id=str(uuid4()), log_level=logging.WARNING)
        )
    assert str(err.value) == "Context must be a type of 'immutables.Map'"


def test_context_has_required_attributes(spp_handler):
    with pytest.raises(ContextError) as err:
        spp_handler.set_context(immutables.Map(my_var="test"))
    assert (
        str(err.value)
        == "Context must contain required arguments: log_correlation_id, log_level"
    )


def test_context_can_be_temporarily_overridden(logger, spp_handler, log_stream):
    spp_handler.set_context(
        immutables.Map(
            log_correlation_id="default_correlation_id",
            log_level=logging.INFO,
        )
    )
    logger.info("my info log message")
    logger.debug("a debug message")
    with spp_handler.override_context(
        immutables.Map(
            log_correlation_id="override_correlation_id",
            log_level=logging.DEBUG,
        )
    ):
        logger.debug("my overridden debug")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 2
    assert log_messages[0]["log_correlation_id"] == "default_correlation_id"
    assert log_messages[0]["description"] == "my info log message"
    assert log_messages[1]["log_correlation_id"] == "override_correlation_id"
    assert log_messages[1]["description"] == "my overridden debug"
