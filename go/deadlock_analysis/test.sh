#!/bin/bash

for i in {1..100000}
do
  1>&2 echo "Welcome $i times"
done

echo 123