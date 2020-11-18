import getpass
import json
import logging
import sys
from datetime import datetime
from typing import IO
from uuid import uuid4

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
        context: dict = None,
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
            context = dict(
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
