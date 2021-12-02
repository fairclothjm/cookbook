# The Makefile `install` and `run` targets will install and run the app in a
# python virtual environment with Python 3's venv module. As long as you have
# Python 3 there should be no other requirements to run the app.
#
# See https://docs.python.org/3/library/venv.html)

.DEFAULT_GOAL := run

install:
	@command -v python3 >/dev/null 2>&1 || { echo >&2 "Python 3 is required"; exit 1; }
	[ -d ".venv" ] || python3 -m venv .venv
	. .venv/bin/activate; pip install -Ur requirements.txt

run:
	. .venv/bin/activate; python3 main.py

clean:
	rm -rf .venv __pycache__


.PHONY: install clean run
