package common_interactor

type DefaultInteractorGateway interface {
	CancelTransaction() error
	CommitTransaction() error
	StartTransaction() error
}
