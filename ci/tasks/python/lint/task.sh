#!/bin/sh

set -eu

pip install poetry

cd spp-logger-git/python

poetry install

poetry run black --check .
poetry run isort --check .
poetry run flake8 --max-line-length=88 .
poetry run mypy --config-file .mypy.ini .
