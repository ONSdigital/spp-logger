import logging
import sys
from typing import IO

from .config import SPPLoggerConfig
from .handler import SPPHandler


class SPPLogger(logging.Logger):
    def __init__(
        self,
        name: str,
        config: SPPLoggerConfig,
        stream: IO = sys.stdout,
    ) -> None:
        super().__init__(name)
        handler = SPPHandler(
            config=config,
            stream=stream,
        )
        self.handlers = [handler]
