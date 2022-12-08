#!/bin/bash

up() {
  goose mysql "$MYSQL_CONNECTION_STRING" up
}
up
ret=$?
while [ $ret -ne 0 ]; do
  up
  ret=$?
  sleep 1
done


