import json
from uuid import UUID


def parse_log_lines(logs):
    log_lines = []
    split_logs = logs.split("\n{")
    for line in split_logs:
        if line.startswith('"'):
            line = "{" + line
        log_lines.append(json.loads(line))
    return log_lines


def is_valid_uuid(uuid_to_test):
    try:
        uuid = UUID(uuid_to_test, version=4)
    except ValueError:
        return False
    return str(uuid) == uuid_to_test


def is_json(log_line):
    try:
        json.loads(log_line)
    except Exception:
        return False
    return True
