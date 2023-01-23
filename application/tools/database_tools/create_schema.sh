#!/bin/bash
function help(){
    echo "Usage: ./create_schema.sh -t <migration name>" 1>&2
    echo "ex)" 1>&2
    echo "./create_schema.sh -t test" 1>&2
    exit 1
}

M_VERSION=""
while [ ${#} -gt 0 ]
    do
        case $1 in
        '-h'|'--help' )
            help
            ;;
        '-t' )
            if [ -z "$2" ]; then
            help
            else
                M_VERSION="$2"
                shift 2
            fi
            ;;
      esac
    done
if [ -z "$M_VERSION" ]; then
    help
else
    go generate ./model/orm/ent
    go run ./tools/database_tools/migration.go "$M_VERSION"
fi

