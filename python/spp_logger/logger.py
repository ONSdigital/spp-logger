import logging
import sys
from contextlib import contextmanager
from typing import IO, Iterator, Union

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
        super().__init__(name, logging.DEBUG)
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

    @property
    def context(self) -> immutables.Map:
        return self.spp_handler.context

    @contextmanager
    def override_context(self, context: immutables.Map) -> Iterator[None]:
        main_context = self.context
        try:
            self.set_context(context)
            yield
        finally:
            self.set_context(main_context)

    def setLevel(self, level: Union[int, str]) -> None:
        raise LogLevelException(
            "SPPLogger does not support setting log level this way. "
            + "Please set the log level using the 'log_level' attribute "
            + "on your context"
        )


class LogLevelException(Exception):
    pass
