#!/bin/bash
###################################################################################################
# A place for all things bash and closely related scripting languages.                            #
###################################################################################################

# Tested on:
#	GNU bash, version 4.4.20(1)-release (x86_64-pc-linux-gnu)
# Description:
#	A simple for loop, with a statement checking if i is 2 or not
#	If i is 2, then echo with the normal echo command, else echo without a new line.
for i in {0..2}; do
  if ((i == 2)); then
    echo "ha"
  else
    echo -n "ha"
  fi
done
