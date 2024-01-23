#!/bin/bash

# Delete captured go-replays
rm -rf requests*

# Delete application logs
rm -rf prod/logs/*
rm -rf qa/logs/*