package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch02/ex02/metricConv"
)

func main() {
	flag.Parse()
	if err := execProc(flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
}

func execProc(args []string) error {
	var ts []float64
	if len(args) > 0 {
		for _, arg := range args {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				return err
			}
			ts = append(ts, t)
		}
	} else {
		//argsのサイズが0なので標準入力を受け取る
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				return err
			}
			ts = append(ts, t)
		}

	}
	printMetricConv(ts)
	return nil
}

func printMetricConv(args []float64) {
	for _, t := range args {
		f := metricConv.Fahrenheit(t)
		c := metricConv.Celsius(t)
		m := metricConv.Meter(t)
		ft := metricConv.Feet(t)
		kg := metricConv.KiloGram(t)
		p := metricConv.Pond(t)
		fmt.Printf("%s = %s, %s = %s,\n"+
			"%s = %s,%s = %s \n"+
			"%s= %s,%s = %s\n",
			f, metricConv.FToC(f), c, metricConv.CToF(c),
			m, metricConv.M2ft(m), ft, metricConv.Ft2m(ft),
			kg, metricConv.Kg2pd(kg), p, metricConv.Pd2kg(p),
		)
	}
}
