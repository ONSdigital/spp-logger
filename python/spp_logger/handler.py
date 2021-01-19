import io
import json
import logging
import sys
import traceback
from contextlib import contextmanager
from datetime import datetime
from types import TracebackType
from typing import IO, Iterator, Optional, Tuple, Union
from uuid import uuid4

import immutables
import pytz as pytz

from .config import SPPLoggerConfig

CONTEXT_REQUIRED_FIELDS = ["log_level", "log_correlation_id"]


class SPPHandler(logging.StreamHandler):
    def __init__(
        self,
        config: SPPLoggerConfig,
        context: immutables.Map = None,
        log_level: Union[int, str] = logging.INFO,
        stream: IO = sys.stdout,
    ) -> None:
        self.config = config
        super().__init__(stream=stream)
        if context is None:
            context = immutables.Map(
                log_correlation_id=str(uuid4()),
                log_correlation_type="AUTO",
                log_level=self.log_level_int(log_level),
            )
        self._context = self.set_context(context)
        self.level = self._context.get("log_level")

    def makeRecord(self, *args, **kwargs):
        return super().makeRecord(*args, **kwargs)

    def format(self, record: logging.LogRecord) -> str:
        log_message = {
            "log_level": record.levelname,
            "timestamp": self.get_timestamp(record),
            "description": record.getMessage(),
            "service": self.config.service,
            "component": self.config.component,
            "environment": self.config.environment,
            "deployment": self.config.deployment,
        }
        if record.exc_info:
            log_message["exception_details"] = self._format_exception_details(record)
        extra = {}
        if hasattr(record, "_extra") and record._extra is not None:  # type: ignore
            extra = record._extra  # type: ignore
        return json.dumps(
            {
                **log_message,
                **extra,
                **{
                    k: self._context[k] for k in self._context if k not in ["log_level"]
                },
                "configured_log_level": self.format_log_level(self.level),
            }
        )

    def _format_exception_details(self, record: logging.LogRecord) -> str:
        s = ""
        if record.exc_info:
            # Cache the traceback text to avoid converting it multiple times
            # (it's constant anyway)
            if not record.exc_text:
                record.exc_text = self._format_exception(record.exc_info)
        if record.exc_text:
            if s[-1:] != "\n":
                s = s + "\n"
            s = s + record.exc_text
        if record.stack_info:
            if s[-1:] != "\n":
                s = s + "\n"
            s = s + record.stack_info
        return s

    def _format_exception(
        self,
        ei: Union[
            Tuple[type, BaseException, Optional[TracebackType]], Tuple[None, None, None]
        ],
    ) -> str:
        sio = io.StringIO()
        tb = ei[2]
        # See issues #9427, #1553375. Commented out for now.
        # if getattr(self, 'fullstack', False):
        #    traceback.print_stack(tb.tb_frame.f_back, file=sio)
        traceback.print_exception(ei[0], ei[1], tb, None, sio)
        s = sio.getvalue()
        sio.close()
        if s[-1:] == "\n":
            s = s[:-1]
        return s

    def get_timestamp(self, record: logging.LogRecord) -> str:
        tz = pytz.timezone(self.config.timezone)
        return datetime.fromtimestamp(record.created, tz).isoformat()

    def set_context_attribute(self, attribute_name: str, attribute_value: str) -> None:
        if attribute_name in self._context:
            raise ImmutableContextError.attribute_error(attribute_name)
        self._context = self._context.set(attribute_name, attribute_value)

    @property
    def context(self) -> immutables.Map:
        return self._context.set(
            "log_level", self.format_log_level(self._context["log_level"])
        )

    def set_context(self, context: immutables.Map) -> immutables.Map:
        if type(context) is not immutables.Map:
            raise ImmutableContextError("Context must be a type of 'immutables.Map'")
        if not all(key in context for key in CONTEXT_REQUIRED_FIELDS):
            raise ContextError(
                "Context must contain required arguments: "
                + ", ".join(CONTEXT_REQUIRED_FIELDS)
            )
        context = context.set("log_level", self.log_level_int(context["log_level"]))
        self._context = context
        self.level = self._context.get("log_level")
        return self._context

    @contextmanager
    def override_context(self, context: immutables.Map) -> Iterator[None]:
        main_context = self._context
        try:
            self.set_context(context)
            yield
        finally:
            self.set_context(main_context)

    def format_log_level(self, log_level: Union[int, str]) -> str:
        if type(log_level) == int:
            return logging.getLevelName(log_level)
        return str(log_level)

    def log_level_int(self, log_level: Union[int, str]) -> int:
        if type(log_level) == str:
            return logging.getLevelName(log_level)
        return int(log_level)


class ImmutableContextError(Exception):
    pass

    @classmethod
    def attribute_error(cls, attribute_name: str) -> "ImmutableContextError":
        return cls(
            "Context attributes are immutable, could not override "
            + f"'{attribute_name}'"
        )


class ContextError(Exception):
    pass
