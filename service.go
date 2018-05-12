package meld

type Service interface {
	CreateAWSEC2() (string, error)
	ReadAWSEC2() (string, error)
	UpdateAWSEC2() (string, error)
	DeleteAWSEC2() (string, error)
}

type service struct {
}

func (s *service) CreateAWSEC2() (string, error) {
	return "Creating a cluster", nil
}

func (s *service) ReadAWSEC2() (string, error) {
	return "Reading a cluster", nil
}

func (s *service) UpdateAWSEC2() (string, error) {
	return "Updating a cluster", nil
}

func (s *service) DeleteAWSEC2() (string, error) {
	return "Deleting a cluster", nil
}

func NewService() Service {
	return &service{}
}
