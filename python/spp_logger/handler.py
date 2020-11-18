import getpass
import json
import logging
import sys
from datetime import datetime
from typing import IO
from uuid import uuid4

import immutables
import pytz as pytz


class SPPHandler(logging.StreamHandler):
    def __init__(
        self,
        service: str,
        component: str,
        environment: str,
        deployment: str,
        user: str = None,
        timezone: str = "UTC",
        context: immutables.Map = None,
        log_level: int = logging.INFO,
        stream: IO = sys.stdout,
    ) -> None:
        self.service = service
        self.component = component
        self.environment = environment
        self.deployment = deployment
        self.user = user
        self.timezone = timezone
        super().__init__(stream=stream)
        if context is None:
            context = immutables.Map(
                log_correlation_id=str(uuid4()),
                log_level_conf=logging.getLevelName(log_level),
            )
        self.context = context

    def format(self, record: logging.LogRecord) -> str:
        log_message = {
            "log_level": record.levelname,
            "timestamp": self.get_timestamp(record),
            "description": record.getMessage(),
            "service": self.service,
            "component": self.component,
            "environment": self.environment,
            "deployment": self.deployment,
            "user": self.get_user(),
        }
        return json.dumps({**log_message, **self.context})

    def get_timestamp(self, record: logging.LogRecord) -> str:
        tz = pytz.timezone(self.timezone)
        return datetime.fromtimestamp(record.created, tz).isoformat()

    def get_user(self) -> str:
        if self.user is None:
            self.user = getpass.getuser()
        return self.user

    def set_context_attribute(self, attribute_name, attribute_value):
        if attribute_name in self.context:
            raise ImmutableContextError(attribute_name)
        self.context = self.context.set(attribute_name, attribute_value)


class ImmutableContextError(Exception):
    def __init__(self, attribute_name: str) -> None:
        self.attribute_name = attribute_name
        super().__init__()

    def __str__(self) -> str:
        return (
            "Context attributes are immutable, could not override "
            + f"'{self.attribute_name}'"
        )
