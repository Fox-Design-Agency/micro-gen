package routeHandlers

// ReturnRouteHandlerLayer will return the byte slice of the route handler layer for the selected service
func ReturnRouteHandlerLayer(microType, name, modelName, projectName string, hasCRUD, hasHelpers bool) ([]byte, error) {

	switch microType {
	case "go":
		// call top
		topString, err := returnGoRouteHandlerTop(name, projectName)
		if err != nil {
			// handle err
		}

		// call methods section
		methodsString, err := returnGoRouteHandlerMethods(name, modelName, hasCRUD, hasHelpers)
		if err != nil {
			// handle err
		}

		return []byte(topString + methodsString), nil

	default:
		return nil, nil
	}
}
