import json
import logging
import sys
import getpass
from datetime import datetime
from typing import IO

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
        stream: IO = sys.stdout,
    ) -> None:
        self.service = service
        self.component = component
        self.environment = environment
        self.deployment = deployment
        self.user = user
        self.timezone = timezone
        super().__init__(stream=stream)

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
        return json.dumps(log_message)

    def get_timestamp(self, record: logging.LogRecord) -> str:
        tz = pytz.timezone(self.timezone)
        return datetime.fromtimestamp(record.created, tz).isoformat()

    def get_user(self) -> str:
        if self.user is None:
            self.user = getpass.getuser()
        return self.user
