package routeHandlers

// ReturnRouteHandlerLayer will return the byte slice of the route handler layer for the selected service
func ReturnRouteHandlerLayer(microType, name, projectName string, hasCRUD bool) ([]byte, error) {

	switch microType {
	case "go":
		// call top
		topString, err := returnGoRouteHandlerTop(name, projectName)
		if err != nil {
			// handle err
		}

		// call methods section
		methodsString, err := returnGoRouteHandlerMethods(name, hasCRUD)
		if err != nil {
			// handle err
		}

		return []byte(topString + methodsString), nil

	default:
		return nil, nil
	}
}
