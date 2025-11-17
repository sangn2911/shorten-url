#!/bin/zsh
export $(grep -v '^#' env/.env | xargs)