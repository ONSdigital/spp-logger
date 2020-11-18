from unittest import mock

from freezegun import freeze_time
from helpers import is_json, parse_log_lines


@freeze_time("2020-11-13")
def test_handler_logs(logger, log_stream):
    logger.info("my info log message")
    log_line = log_stream.getvalue()
    assert (
        is_json(log_line) is True
    ), f"Expected log lines to be JSON but was: '{log_line}'"
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 1
    assert len(log_messages[0]) == 8
    assert log_messages[0]["log_level"] == "INFO"
    assert log_messages[0]["timestamp"] == "2020-11-13T00:00:00+00:00"
    assert log_messages[0]["description"] == "my info log message"
    assert log_messages[0]["service"] == "test-service"
    assert log_messages[0]["component"] == "test-component"
    assert log_messages[0]["environment"] == "dev"
    assert log_messages[0]["deployment"] == "test-deployment"
    assert log_messages[0]["user"] == "test-user"


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
    spp_handler.user = None
    assert spp_handler.get_user() == "my_test_user"
    spp_handler.get_user()
    spp_handler.get_user()
    mock_get_user.assert_called_once()
