package server

const LogLevelInfo = 0
const LogLevelWarning = 1
const LogLevelError = 2
const LogLevelDebug = 3

type OptFunc func(*Opts)
type Opts struct {
	listenAddr string
	isSSL      bool
	logLevel   int
}
type Server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		listenAddr: ":8080",
		isSSL:      false,
		logLevel:   LogLevelInfo,
	}
}

func WithListenAddr(addr string) OptFunc {
	return func(o *Opts) {
		o.listenAddr = addr
	}
}
func WithSSL(o *Opts) {
	o.isSSL = true
}
func WithLogLevel(level int) OptFunc {
	return func(o *Opts) {
		o.logLevel = level
	}
}

func NewServer(optFuncs ...OptFunc) Server {
	o := defaultOpts()
	for _, optFunc := range optFuncs {
		optFunc(&o)
	}
	return Server{
		o,
	}
}

func (s Server) HandleSomeStuff() {
	println(s.listenAddr, s.isSSL, s.logLevel)
}
