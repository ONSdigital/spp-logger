import logging
import sys
from typing import IO

from spp_logger import SPPHandler


class SPPLogger(logging.Logger):
    def __init__(
        self,
        name: str,
        service: str,
        component: str,
        environment: str,
        deployment: str,
        user: str = None,
        timezone: str = "UTC",
        stream: IO = sys.stdout,
    ) -> None:
        super().__init__(name)
        handler = SPPHandler(
            service=service,
            component=component,
            environment=environment,
            deployment=deployment,
            user=user,
            timezone=timezone,
            stream=stream,
        )
        self.handlers = [handler]
