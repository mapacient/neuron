package edge

type Edge interface {
	Attach() (addr uintptr, err error)
	Detach() error
	Close() error
}

const (
	EDGE_CREAT  = 0x01
	EDGE_RDONLY = 0x02
	EDGE_WRONLY = 0x04
	EDGE_EXCL   = 0x08
	EDGE_RDWR   = EDGE_RDONLY | EDGE_WRONLY
)

//Open open call
func Open(name string, size uint64, mode int) (Edge, error) {

	return OpenSharedMemory(name, size, mode)
}
