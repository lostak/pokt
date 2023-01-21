package keeper

import (
	"flag"
)

var (
	port = flag.Int("port", 8080, "Tx server port")
)

type PoktServer struct {
	UnimplementedMsgServer
}
