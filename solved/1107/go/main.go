package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// var n
	// fmt.Fscanf(in, "%d", )

	// blks = make([], 0)
}

func findNearest(n int, blocked string) int {
	mi := n

	for ; mi >= 0; mi-- {
		ma := strconv.Itoa(mi)

		if !strings.ContainsAny(ma, blocked) {
			break
		}
	}

	mx := n

	for ; ; mx++ {
		ma := strconv.Itoa(mx)

		if !strings.ContainsAny(ma, blocked) {
			break
		}
	}

	if mi >= 0 && n-mi < mx-n {
		return mi
	} else {
		return mx
	}
}
