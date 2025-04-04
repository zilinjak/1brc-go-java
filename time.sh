#!/bin/zsh


echo "Baseline: " >> times.txt

{ /opt/homebrew/opt/gnu-time/libexec/gnubin/time -v ./calculate_average_baseline.sh } > result.txt 2>>times.txt
echo "Go: " >> times.txt
go build 1brc.go
chmod +x 1brc
{ /opt/homebrew/opt/gnu-time/libexec/gnubin/time -v ./1brc } > result_go.txt 2>>times.txt
rm -rf 1brc

cat times.txt

