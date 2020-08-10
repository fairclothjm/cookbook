#!/usr/bin/env python3
"""
    argparse usage code snippet
"""
import argparse

def setup_argparse():
    """
    :returns: argument parser object
    :rtype: argparse.ArgumentParser
    """
    p = argparse.ArgumentParser(description="argparse example", prog="my-cli")

    p.add_argument(
        "-n",
        "--name",
        type=str,
        help="the name",
        default="default-name",
        required=True,
    )

    p.add_argument(
        "-c",
        "--count",
        type=int,
        help="the count",
        default=0,
    )

    p.add_argument(
        "-f",
        "--fork",
        type=bool,
        help="fork as a daemon process",
        default=False,
    )

    p.add_argument(
        "-i",
        "--interval",
        type=int,
        help="space-separated list of time intervals",
        nargs="+",
    )

    return p


if __name__ == "__main__":
    parser = setup_argparse()

    args = parser.parse_args()
    print(args)
