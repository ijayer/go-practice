package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define Usage
	flag.Usage = func() {
		fmt.Println("Usage: Test Command line.")
		flag.PrintDefaults()
	}

	// Subcommands
	countCmd := flag.NewFlagSet("count", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Count subcommand
	countTextPtr := countCmd.String("text", "", "    Text to parse. (Required)")
	countMetricPtr := countCmd.String("metric", "chars", "    Metric {chars|words|lines|substring}. (Required)")
	countSubstringPtr := countCmd.String("substring", "", "    The substring to be counted. Required for --metric=substring")
	countUniquePtr := countCmd.Bool("unique", false, "    Measure unique values of a metric.")

	// List subcommand flag pointers
	listTextPtr := listCmd.String("text", "", "    Text to parse. (Required)")
	listMetricPtr := listCmd.String("metric", "chars", "    Metric <chars|words|lines>. (Required)")
	listUniquePtr := listCmd.Bool("unique", false, "    Measure unique values of a metric.")

	// Verify that a subcommand has been provided
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Switch on the subcommands
	// fmt.Printf("os.Args: %v", os.Args)
	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
	case "count":
		countCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed functions.
	if listCmd.Parsed() {
		if *listTextPtr == "" {
			listCmd.PrintDefaults()
			os.Exit(1)
		}

		metricChoices := map[string]bool{"chars": true, "words": true, "lines": true}
		if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
			listCmd.PrintDefaults()
			os.Exit(1)
		}
	}
	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *listTextPtr, *listMetricPtr, *listUniquePtr)

	if countCmd.Parsed() {
		// Required Flags
		if *countTextPtr == "" {
			countCmd.PrintDefaults()
			os.Exit(1)
		}
		// If the metric flag is substring, the substring flag is required
		if *countMetricPtr == "substring" && *countSubstringPtr == "" {
			countCmd.PrintDefaults()
			os.Exit(1)
		}
		//If the metric flag is not substring, the substring flag must not be used
		if *countMetricPtr != "substring" && *countSubstringPtr != "" {
			fmt.Println("--substring may only be used with --metric=substring.")
			countCmd.PrintDefaults()
			os.Exit(1)
		}
		//Choice flag
		metricChoices := map[string]bool{"chars": true, "words": true, "lines": true, "substring": true}
		if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
			countCmd.PrintDefaults()
			os.Exit(1)
		}
		//Print
		fmt.Printf("textPtr: %s, metricPtr: %s, substringPtr: %v, uniquePtr: %t\n", *countTextPtr, *countMetricPtr, *countSubstringPtr, *countUniquePtr)
	}
}
