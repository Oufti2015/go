#!/bin/bash

echo Publishing $1 as Lesson $2

echo -----------------------
echo Add $1
git add $1

echo -----------------------
echo Commit "Lesson $2"
git commit -m "Lesson $2"

echo -----------------------
echo Push origin
git push origin

exit 0