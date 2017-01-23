#!/usr/bin/env python3
"""Convert to titlecase."""
from sys import stdin, stdout

from titlecase import titlecase


if __name__ == "__main__":
    if stdin.isatty():
        try:
            while True:
                print(titlecase(input("> ")))
        except (KeyboardInterrupt, EOFError):
            print("\nGood Bye")
    else:
        stdout.write(titlecase(stdin.read()))
