import os
from dataclasses import dataclass


@dataclass
class SPPLoggerConfig:
    service: str
    component: str
    environment: str
    deployment: str
    timezone: str = "UTC"

    @classmethod
    def from_env(cls) -> "SPPLoggerConfig":
        return cls(
            service=os.environ["SPP_SERVICE"],
            component=os.environ["SPP_COMPONENT"],
            environment=os.environ["SPP_ENVIRONMENT"],
            deployment=os.environ["SPP_DEPLOYMENT"],
            timezone=os.getenv("TIMEZONE", "UTC"),
        )
