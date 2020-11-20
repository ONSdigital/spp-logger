#!/bin/sh

set -eu

pip install poetry

cd spp-logger-git/python

poetry install

poetry run python -m pytest -p no:warnings
