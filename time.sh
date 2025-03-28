#!/bin/zsh


echo "Baseline: " > times.txt
{ time ./calculate_average_baseline.sh } > result.txt 2>>times.txt
echo "Go: " >> times.txt
{ time go run 1brc.go } > result_go.txt 2>>times.txt

cat times.txt

