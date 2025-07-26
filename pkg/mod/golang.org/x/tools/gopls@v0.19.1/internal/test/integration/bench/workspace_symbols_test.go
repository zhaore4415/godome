// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bench

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

var symbolQuery = flag.String("symbol_query", "test", "symbol query to use in benchmark")

// BenchmarkWorkspaceSymbols benchmarks the time to execute a workspace symbols
// request (controlled by the -symbol_query flag).
func BenchmarkWorkspaceSymbols(b *testing.B) {
	for name := range repos {
		b.Run(name, func(b *testing.B) {
			env := getRepo(b, name).sharedEnv(b)
			start := time.Now()
			symbols := env.Symbol(*symbolQuery) // warm the cache

			if testing.Verbose() {
				fmt.Printf("Results (after %s):\n", time.Since(start))
				for i, symbol := range symbols {
					fmt.Printf("\t%d. %s (%s)\n", i, symbol.Name, symbol.ContainerName)
				}
			}

			b.ResetTimer()

			if stopAndRecord := startProfileIfSupported(b, env, qualifiedName(name, "workspaceSymbols")); stopAndRecord != nil {
				defer stopAndRecord()
			}

			for b.Loop() {
				env.Symbol(*symbolQuery)
			}
		})
	}
}
