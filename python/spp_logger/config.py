import os
from dataclasses import dataclass
from typing import Optional


@dataclass
class SPPLoggerConfig:
    service: str
    component: str
    environment: str
    deployment: str
    user: Optional[str] = None
    timezone: str = "UTC"

    @classmethod
    def from_env(cls) -> "SPPLoggerConfig":
        return cls(
            service=os.environ["SPP_SERVICE"],
            component=os.environ["SPP_COMPONENT"],
            environment=os.environ["SPP_ENVIRONMENT"],
            deployment=os.environ["SPP_DEPLOYMENT"],
            user=os.getenv("SPP_USER", None),
            timezone=os.getenv("TIMEZONE", "UTC"),
        )
