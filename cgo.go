package goprisma

/*
#cgo darwin,!arm LDFLAGS: -L"${SRCDIR}/lib/darwin" -Wl,-rpath,"${SRCDIR}/lib/darwin" -lquery_engine_c_api
#cgo darwin,arm LDFLAGS: -L"${SRCDIR}/lib/darwin-aarch64" -Wl,-rpath,"${SRCDIR}/lib/darwin" -lquery_engine_c_api
#cgo linux LDFLAGS: -L"${SRCDIR}/lib/linux" -lquery_engine_c_api
#cgo windows LDFLAGS: -L"${SRCDIR}/lib/windows" -lquery_engine_c_api
*/
import "C"
