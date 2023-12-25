#!/bin/sh

if [ $# -ne 2 ]; then
  echo "引数がおかしい"
  exit 9
fi

case $1 in
  0 ) echo "migrationファイルを作成"
      migrate create -ext sql -dir infra/database/migrations -seq $2
      ;;
  1 ) echo "migrationファイルを適応"
      migrate --path infra/database/migrations --database 'mysql://docker:docker@tcp(db:3306)/bowling_score' -verbose $2
      ;;
  * ) echo "実行引数がおかしい"
      exit 9
      ;;
esac