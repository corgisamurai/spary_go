#!/bin/bash

echo "hoge"
mysql -u root -p < ../db/drop_database.sql
mysql -u root -p < ../db/init.sql
mysql -u root -p < ../db/create_table.sql


