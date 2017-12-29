package debug

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/ckeyer/commons/version"
)

const debugInfo = `
version      - current binary version
goroutine    - stack traces of all current goroutines
heap         - a sampling of all heap allocations
threadcreate - stack traces that led to the creation of new OS threads
block        - stack traces that led to blocking on synchronization primitives
`

func init() {
	start := time.Now()
	chDebug := make(chan os.Signal, 1)
	signal.Notify(chDebug, syscall.SIGUSR1)

	go func() {
		for s := range chDebug {
			switch s {
			case syscall.SIGUSR1:
				f, err := os.OpenFile("/tmp/stack-trace.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
				if err != nil {
					break
				}
				fmt.Fprint(f, debugInfo)
				fmt.Fprint(f, "\n\nVERSION\n\n")
				fmt.Fprintf(f, "version ----------- %s\n", version.GetCompleteVersion())
				fmt.Fprintf(f, "git commit -------- %s\n", version.GetGitCommit())
				fmt.Fprintf(f, "build at: --------- %s\n", version.GetBuildAt())
				fmt.Fprintf(f, "start at: --------- %s\n", start)
				fmt.Fprint(f, "\n\nGOROUTINE\n\n")
				pprof.Lookup("goroutine").WriteTo(f, 2)
				fmt.Fprint(f, "\n\nHEAP\n\n")
				pprof.Lookup("heap").WriteTo(f, 1)
				fmt.Fprint(f, "\n\nTHREADCREATE\n\n")
				pprof.Lookup("threadcreate").WriteTo(f, 1)
				fmt.Fprint(f, "\n\nBLOCK\n\n")
				pprof.Lookup("block").WriteTo(f, 1)
				f.Close()
			}
		}
	}()
}
