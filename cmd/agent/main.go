package main
import "fmt"
type Watcher struct { namespace string; configs []string }
func (w *Watcher) Watch() { fmt.Printf("Watching %s in %s\n", w.configs, w.namespace) }
func main() { w := &Watcher{namespace: "default", configs: []string{"app-config"}}; w.Watch() }
