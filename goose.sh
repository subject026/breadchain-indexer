#! /usr/bin/bash

cd sql/schema

goose postgres postgres://root:password@localhost:5432/breadchain-indexer $1