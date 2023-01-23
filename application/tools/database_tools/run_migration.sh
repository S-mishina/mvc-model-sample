#!/bin/bash
function help(){
    echo "Usage: ./run_migration.sh -v <migration version> -e <mysql url> -f <migration file (default:'migration/')>" 1>&2
    echo "ex)" 1>&2
    echo "./run_migration.sh -v 20230110190309 -e root:pass@localhost:3306/test -f migration/" 1>&2
    exit 1
}
VERSION=""
CONFIG=""
FILE="migration/"
while [ ${#} -gt 0 ]
    do
        case $1 in
        '-h'|'--help' )
            help
            ;;
        '-v' )
            if [ -z "$2" ]; then
            help
            else
                VERSION="$2"
                shift 2
            fi
            ;;
        '-e' )
            if [ -z "$2" ]; then
                help
            else
                CONFIG="$2"
                shift 2
            fi
            ;;
        '-f' )
            if [ -z "$2" ]; then
                help
            else
                FILE="$2"
                shift 2
            fi
          ;;
      esac
    done
echo $FILE
if [ -z "$VERSION" ] || [ -z "$CONFIG" ]; then
    help
else
    /usr/local/bin/atlas migrate apply --dir "file://migration/" --var test="$VERSION" --url mysql://"$CONFIG"
fi