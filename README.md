# Spary Go
Sapry backend  
No SPA, No life! You're making more better SPA life.

## Tech
|  Name  |  version  |
| ---- | ---- |
| Golang | 1.7.4 | 
| ---- | ---- |

## Init
$ mysql -u root -p < db/init.sql  
$ mysql -u root -p < db/create_table.sql

## Batch
$ Import Spa Info : go run main.go ImportOnsenList

## Test
setting ROOT_PATH at .env
$ sh test.sh
$ sh watch_test.sh
dependency: brew install fswatch

