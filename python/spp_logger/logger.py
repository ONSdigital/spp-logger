import logging
import sys
from contextlib import contextmanager
from types import TracebackType
from typing import IO, Any, Iterator, Mapping, Optional, Tuple, Union

import immutables

from .config import SPPLoggerConfig
from .handler import SPPHandler


class SPPLogger(logging.Logger):
    def __init__(
        self,
        name: str,
        config: SPPLoggerConfig,
        context: immutables.Map = None,
        stream: IO = sys.stdout,
    ) -> None:
        super().__init__(name, logging.DEBUG)
        handler = SPPHandler(
            config=config,
            context=context,
            stream=stream,
        )
        self.spp_handler = handler
        self.handlers = [handler]

    def set_context_attribute(self, attribute_name: str, attribute_value: str) -> None:
        self.spp_handler.set_context_attribute(attribute_name, attribute_value)

    def set_context(self, context: immutables.Map) -> immutables.Map:
        return self.spp_handler.set_context(context)

    def makeRecord(
        self,
        name: str,
        level: int,
        fn: str,
        lno: int,
        msg: str,
        args: Union[Tuple[Any, ...], Mapping[str, Any]],
        exc_info: Union[
            Tuple[type, BaseException, Optional[TracebackType]],
            Tuple[None, None, None],
            None,
        ],
        func: Optional[str] = None,
        extra: Optional[Mapping[str, Any]] = None,
        sinfo: Optional[str] = None,
    ) -> logging.LogRecord:
        record = super().makeRecord(
            name, level, fn, lno, msg, args, exc_info, func, extra, sinfo
        )
        record._extra = extra  # type: ignore
        return record

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
