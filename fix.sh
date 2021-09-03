#!/bin/bash
find . -type f -name '*.go' -exec sed -e 's#github.com/filebrowser/filebrowser#github.com/dolfly/filebrowser#g' {} \;
