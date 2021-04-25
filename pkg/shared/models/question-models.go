package models

// Questions is the stuct type to hold
// the question data to be utilized during
// the build
type Questions struct {
	ProjectName string
	HasHelpers  bool
	// IsCLI       bool
	HasDB       bool
	SubServices []*SubService
	Port        string
}

// SubService is the struct type to hold subservice
// information to generate the desired subservices
type SubService struct {
	SubServiceName  string
	HasDB           bool
	HasCRUD         bool
	HasRouteHandler bool
	ModelName       string
}
