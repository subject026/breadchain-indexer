#! /usr/bin/bash

cd sql/schema

goose postgres postgres://root:password@localhost:5432/breadchain_indexer $1
