#!/bin/bash

echo cd json-std
cd json-std
go test -bench=.

echo cd ../json-iterator
cd ../json-iterator
go test -bench=.
