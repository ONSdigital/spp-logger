from helpers import parse_log_lines


def test_logger(spp_logger, log_stream):
    spp_logger.info("my info log message")
    log_messages = parse_log_lines(log_stream.getvalue())
    assert len(log_messages) == 1
    assert log_messages[0]["description"] == "my info log message"
