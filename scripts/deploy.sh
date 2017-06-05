#!/bin/bash

set -e

# The deploy stage within Travis-CI dirtys the working tree. This will cause
# the docker tagging to fail. Reset to HEAD as a workaround.
git reset --hard HEAD
conform enforce deploy
