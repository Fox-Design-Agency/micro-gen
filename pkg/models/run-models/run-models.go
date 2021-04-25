package run

import (
	"log"
)

// ReturnMainFile will return the string for the main.go file
func ReturnMainFile(microType string, serviceArray, routeHandlerArray []string, hasDB bool, projectName string) ([]byte, error) {
	switch microType {
	case "go":
		// top + init
		topString, err := returnGoMainTopSection(projectName)
		if err != nil {
			// handle err
			log.Println(err)
		}
		// main configuration
		configString, err := returnGoMainConfigurationSection(hasDB, serviceArray)
		if err != nil {
			// handle err
			log.Println(err)
		}
		// initialize router and router handlers
		routerString, err := returnGoMainRouterSection(hasDB, routeHandlerArray)
		if err != nil {
			// handle err
			log.Println(err)
		}
		// initialize middleware
		middlewareString, err := returnGoMainMiddlewareSection()
		if err != nil {
			// handle err
			log.Println(err)
		}
		// health check
		healthString, err := returnGoMainHealthSection()
		if err != nil {
			// handle err
			log.Println(err)
		}
		// routes
		routesString, err := returnGoMainRoutesSection(hasDB, routeHandlerArray)
		if err != nil {
			// handle err
			log.Println(err)
		}
		// server
		serverString, err := returnGoMainServerSection()
		if err != nil {
			// handle err
			log.Println(err)
		}

		return []byte(topString +
			configString +
			routerString +
			middlewareString +
			healthString +
			routesString +
			serverString), nil

	default:
		return nil, nil
	}
}

// ReturnConfigFile will return the string for the config.go file
func ReturnConfigFile(microType string) ([]byte, error) {
	switch microType {
	case "go":
		// structs
		structString, err := returnGoConfigStructsSection()
		if err != nil {
			// handle err
			log.Println(err)
		}
		// funcs
		funcString, err := returnGoConfigFuncsSection()
		if err != nil {
			// handle err
			log.Println(err)
		}

		return []byte(structString + funcString), nil

	default:
		return nil, nil
	}
}

// ReturnSecretStuffFile will return the string for the secretStuff.go file
func ReturnSecretStuffFile(microType string) ([]byte, error) {
	switch microType {
	case "go":
		secretString, err := returnGoSetTheSecretStuff()
		if err != nil {
			// handle err
			log.Println(err)
		}

		return []byte(secretString), nil
	default:
		return nil, nil
	}

}
