import immutables

from spp_logger import context_to_dict, dict_to_context


def test_context_to_dict():
    context = immutables.Map(
        correlation_id="test-correlation-id",
        log_level="INFO",
    )
    assert context_to_dict(context) == {
        "correlation_id": "test-correlation-id",
        "log_level": "INFO",
    }


def test_dict_to_context():
    context_dict = {"correlation_id": "test-correlation-id", "log_level": "INFO"}
    assert dict_to_context(context_dict) == immutables.Map(
        correlation_id="test-correlation-id",
        log_level="INFO",
    )
