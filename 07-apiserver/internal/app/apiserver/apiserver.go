package apiserver

type APIServer struct{}

func New() *APIServer {
	return &APIServer{}
}

func (a *APIServer) Start() error {
	return nil
}
