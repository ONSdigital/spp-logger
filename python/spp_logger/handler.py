import logging
import sys
from typing import IO


class SPPHandler(logging.StreamHandler):
    def __init__(
        self,
        service: str,
        component: str,
        environment: str,
        deployment: str,
        user: str = None,
        timezone: str = "UTC",
        stream: IO = sys.stdout,
    ) -> None:
        super().__init__(stream=stream)
