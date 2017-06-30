#!/bin/bash

cd ..
go install
cd scripts

sudo cp ./pac.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable pac.service

