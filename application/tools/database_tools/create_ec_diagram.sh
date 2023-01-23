#!/bin/bash
go run github.com/hedwigz/entviz/cmd/entviz ./model/orm/ent/schema
mv schema-viz.html docs/er_diagram/schema-viz.html
