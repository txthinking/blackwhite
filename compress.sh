#!/bin/bash

./node_modules/.bin/uglifyjs --compress --mangle -- white.pac > _
mv _ white.pac

./node_modules/.bin/uglifyjs --compress --mangle -- black.pac > _
mv _ black.pac

./node_modules/.bin/uglifyjs --compress --mangle -- all.pac > _
mv _ all.pac
