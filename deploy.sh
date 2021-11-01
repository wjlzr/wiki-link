#!/bin/bash
git pull
docker-compose build server workers schedule
docker-compose up -d