package contract

import "net/http"

const KernelKey = "wm:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}
