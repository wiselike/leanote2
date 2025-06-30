#!/bin/sh

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/leanote-of-unofficial" -importPath github.com/wiselike/leanote-of-unofficial -srcPath "$SCRIPTPATH/src" -runMode dev
