#!/bin/bash

year=$1
day=$2

mkdir -p $year/$day
cp -r template/* $year/$day/
open https://adventofcode.com/$year/day/$day
open https://adventofcode.com/$year/day/$day/input
code $year/$day/*
cd $year/$day