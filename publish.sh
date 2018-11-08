#!/bin/bash

echo Publishing $1 as Lesson $2

git add $1

git commit -m "Lesson $2"

git push origin

exit 0