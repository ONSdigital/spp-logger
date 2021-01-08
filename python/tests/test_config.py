import os
from unittest import mock

from spp_logger import SPPLoggerConfig


@mock.patch.dict(
    os.environ,
    {
        "SPP_SERVICE": "test-service",
        "SPP_COMPONENT": "test-component",
        "SPP_ENVIRONMENT": "test-environment",
        "SPP_DEPLOYMENT": "test-deployment",
        "TIMEZONE": "GMT",
    },
)
def test_from_env():
    spp_config = SPPLoggerConfig.from_env()
    assert spp_config.service == "test-service"
    assert spp_config.component == "test-component"
    assert spp_config.environment == "test-environment"
    assert spp_config.deployment == "test-deployment"
    assert spp_config.timezone == "GMT"


@mock.patch.dict(
    os.environ,
    {
        "SPP_SERVICE": "test-service",
        "SPP_COMPONENT": "test-component",
        "SPP_ENVIRONMENT": "test-environment",
        "SPP_DEPLOYMENT": "test-deployment",
    },
)
def test_from_env_has_defaults():
    spp_config = SPPLoggerConfig.from_env()
    assert spp_config.service == "test-service"
    assert spp_config.component == "test-component"
    assert spp_config.environment == "test-environment"
    assert spp_config.deployment == "test-deployment"
    assert spp_config.timezone == "UTC"


def test_init():
    spp_config = SPPLoggerConfig(
        service="test-service",
        component="test-component",
        environment="test-environment",
        deployment="test-deployment",
        timezone="GMT",
    )
    assert spp_config.service == "test-service"
    assert spp_config.component == "test-component"
    assert spp_config.environment == "test-environment"
    assert spp_config.deployment == "test-deployment"
    assert spp_config.timezone == "GMT"


def test_init_has_defaults():
    spp_config = SPPLoggerConfig(
        service="test-service",
        component="test-component",
        environment="test-environment",
        deployment="test-deployment",
    )
    assert spp_config.service == "test-service"
    assert spp_config.component == "test-component"
    assert spp_config.environment == "test-environment"
    assert spp_config.deployment == "test-deployment"
    assert spp_config.timezone == "UTC"
