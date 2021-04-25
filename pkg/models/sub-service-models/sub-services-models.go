package subservices

import (
	"log"
)

// ReturnSubServiceLayer will return the byte slice of the generic SubService
// layer file
func ReturnSubServiceLayer(microType, name, projectName string, hasDB bool) ([]byte, error) {
	switch microType {
	case "go":
		// top section
		topString, err := returnGoSubServiceTop(name, projectName, hasDB)
		if err != nil {
			// handle err
			log.Println(err)
		}
		// definition section
		definitionString, err := returnGoSubServiceSvcDefinition(name, hasDB)
		if err != nil {
			// handle err
			log.Println(err)
		}

		return []byte(topString + definitionString), nil
	default:
		return nil, nil
	}

}

// ReturnSubServiceServiceFile will return the string for the services.go file
func ReturnSubServiceServiceFile(serviceArray []string, hasDB bool) ([]byte, error) {
	// top section
	topString, err := returnServiceFileTopSection()
	if err != nil {
		// handle err
		log.Println(err)
	}
	// with section
	withString, err := returnServiceFileWithSection(hasDB)
	if err != nil {
		// handle err
		log.Println(err)
	}
	// newServicesSection
	newServicesString, err := returnServiceFileNewServiceSection(hasDB, serviceArray)
	if err != nil {
		// handle err
		log.Println(err)
	}
	// Services struct
	servicesStructString, err := returnServiceFileServicesStructSection(hasDB, serviceArray)
	if err != nil {
		// handle err
		log.Println(err)
	}
	// global section
	globalsString, err := returnServiceFileGlobalSection(hasDB)
	if err != nil {
		// handle err
		log.Println(err)
	}

	return []byte(topString +
		withString +
		newServicesString +
		servicesStructString +
		globalsString), nil
}
