import sys

from setuptools import setup

assert sys.version_info >= (3, 6, 0), "spp_logger requires Python 3.6+"
from pathlib import Path  # noqa E402

CURRENT_DIR = Path(__file__).parent
sys.path.insert(0, str(CURRENT_DIR))  # for setuptools.build_meta

setup(
    name="spp_logger",
    description="A logging library to standardise log formats in JSON format",
    url="https://github.com/ONSdigital/spp-logger",
    license="MIT",
    packages=["spp_logger"],
    package_dir={"": "."},
    package_data={"spp_logger": ["py.typed"]},
    python_requires=">=3.6",
    install_requires=[
        "pytz>=2020.4",
    ],
    test_suite="tests",
    classifiers=[
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3 :: Only",
        "Topic :: Software Development :: Libraries :: Python Modules",
    ],
)
