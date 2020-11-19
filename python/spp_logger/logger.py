import logging
import sys
from typing import IO

import immutables

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
        self.spp_handler = handler
        self.handlers = [handler]

    def set_context_attribute(self, attribute_name: str, attribute_value: str) -> None:
        self.spp_handler.set_context_attribute(attribute_name, attribute_value)

    def set_context(self, context: immutables.Map) -> immutables.Map:
        return self.spp_handler.set_context(context)
