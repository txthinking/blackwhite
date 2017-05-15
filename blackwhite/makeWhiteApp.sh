#!/bin/bash

echo 'package blackwhite' > _
echo 'var was string = `' >> _
cat ../white_apps.list >> _
echo '`' >> _
mv _ white_app.go
