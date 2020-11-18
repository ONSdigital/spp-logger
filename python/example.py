import logging
import sys

handler = logging.StreamHandler(stream=sys.stdout)
logger = logging.getLogger()
logger.addHandler(handler)
logger.setLevel(logging.INFO)
logger.info("My info log message")
logger.warning("foobar")
