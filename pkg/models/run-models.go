package models

import (
	"fmt"
	"log"
	"strings"
)

// ReturnMainFile will return the string for the main.go file
func ReturnMainFile(serviceArray []string, hasDB bool, projectName string) ([]byte, error) {
	// top + init
	topString, err := returnMainTopSection(projectName)
	if err != nil {
		// handle err
		log.Println(err)
	}
	// main configuration
	configString, err := returnMainConfigurationSection(hasDB, serviceArray)
	if err != nil {
		// handle err
		log.Println(err)
	}
	// initialize router and router handlers
	routerString, err := returnMainRouterSection(hasDB, serviceArray)
	if err != nil {
		// handle err
		log.Println(err)
	}
	// initialize middleware
	middlewareString, err := returnMainMiddlewareSection()
	if err != nil {
		// handle err
		log.Println(err)
	}
	// health check
	healthString, err := returnMainHealthSection()
	if err != nil {
		// handle err
		log.Println(err)
	}
	// routes
	routesString, err := returnMainRoutesSection(hasDB, serviceArray)
	if err != nil {
		// handle err
		log.Println(err)
	}
	// server
	serverString, err := returnMainServerSection()
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
}

// returnMainTopSection will return the string for the package, imports
// and init section of the main.go file
func returnMainTopSection(projectName string) (string, error) {
	topString := fmt.Sprintf("package main\n\n") +
		fmt.Sprintf("import (\n") +
		fmt.Sprintf("	routeHandlers \"%s/pkg/route-handlers\"", projectName) +
		fmt.Sprintf("	services \"%s/pkg/sub-services\"\n\n", projectName) +
		fmt.Sprintf("	\"fmt\"\n") +
		fmt.Sprintf("	\"log\"\n") +
		fmt.Sprintf("	\"net/http\"\n") +
		fmt.Sprintf("	\"time\"\n\n") +
		fmt.Sprintf("	\"github.com/gorilla/handlers\"\n") +
		fmt.Sprintf("	\"github.com/gorilla/mux\"\n)\n\n")

	initString := fmt.Sprintf("func init() {\n") +
		fmt.Sprintf("// run the secrets, hidden with the .gitignore\n") +
		fmt.Sprintf("// basically just sets the ENV vars for:\n") +
		fmt.Sprintf("// DBHOST, DBNAME, DBUSER, DBPASS, PORT, ENVIRONMENT\n") +
		fmt.Sprintf("setTheSecrets()\n}\n\n")

	return topString + initString, nil
}

// returnMainConfigurationSection will return the string for configuration
// section of the main func in the main.go file
func returnMainConfigurationSection(hasDB bool, serviceArray []string) (string, error) {
	topString := fmt.Sprintf("func main() {\n\n") +
		topCommentBlock +
		fmt.Sprintf("\n/	Configuration\n") +
		bottomCommentBlock +
		fmt.Sprintf("\n\n // load application configuration\n") +
		fmt.Sprintf("cfg := LoadConfig()\n\n") +
		fmt.Sprintf("// default application port\n") +
		fmt.Sprintf("port := cfg.Port\n\n") +
		fmt.Sprintf("// set db config to pass to service intialization\n") +
		fmt.Sprintf("dbCfg := cfg.Database\n\n")

	serviceDeclarationString := fmt.Sprintf("srvcs, err := services.NewServices(\n")
	if hasDB {
		serviceDeclarationString += fmt.Sprintf("services.WithPostgres(dbCfg.Dialect(), dbCfg.Connection()),\n")
	}
	for _, v := range serviceArray {
		serviceDeclarationString += fmt.Sprintf("services.With%s(),\n", strings.Title(v))
	}
	serviceDeclarationString += fmt.Sprintf(")\n") +
		fmt.Sprintf("if err != nil {\n") +
		fmt.Sprintf("// services couldnt be made, should probably panic?\n") +
		fmt.Sprintf("// something here\n") +
		fmt.Sprintf("log.Println(err)\n") +
		fmt.Sprintf("return\n}\n")
	if hasDB {
		serviceDeclarationString += fmt.Sprintf("defer srvcs.Close()\n") +
			fmt.Sprintf("//run migrates\n") +
			fmt.Sprintf("err = srvcs.MigrateDBUP()\n") +
			fmt.Sprintf("if err != nil {\n") +
			fmt.Sprintf("// migrations couldn't happen\n") +
			fmt.Sprintf("log.Println(err)\n}\n\n")
	}

	return topString + serviceDeclarationString, nil
}

// returnMainRouterSection will return the string for the router
// section of the main.go file
func returnMainRouterSection(hasDB bool, serviceArray []string) (string, error) {
	topString := topCommentBlock +
		fmt.Sprintf("\n/	Initialize router and controllers\n") +
		bottomCommentBlock +
		fmt.Sprintf("\n\nr := mux.NewRouter()\n\n")
	for _, v := range serviceArray {
		topString += fmt.Sprintf("%sRH := routeHandlers.New%s(\n", strings.ToLower(v), strings.Title(v)) +
			fmt.Sprintf("	srvcs.%s)\n", strings.Title(v))
	}
	topString += fmt.Sprintf("\n")

	return topString, nil
}

// returnMainMiddlewareSection will return the string for the middleware
// section of the main.go file
func returnMainMiddlewareSection() (string, error) {
	topString := topCommentBlock +
		fmt.Sprintf("\n/	Middleware\n") +
		bottomCommentBlock +
		fmt.Sprintf("\n\n")

	return topString, nil
}

// returnMainHealthSection will return the string for the health
// check section of the main.go file
func returnMainHealthSection() (string, error) {
	topString := topCommentBlock +
		fmt.Sprintf("\n/	Health Check & Container routes\n") +
		bottomCommentBlock +
		fmt.Sprintf("\n\nr.HandleFunc(\"/health\", func(rw http.ResponseWriter, r *http.Request) { rw.WriteHeader(http.StatusOK) })\n\n")

	return topString, nil
}

// returnMainRoutesSection will return the string for the routes
// section of the main.go file
func returnMainRoutesSection(hasDB bool, serviceArray []string) (string, error) {
	topString := ""
	for _, v := range serviceArray {
		topString += topCommentBlock +
			fmt.Sprintf("\n/	%s routes\n", strings.Title(v)) +
			bottomCommentBlock +
			fmt.Sprintf("\n")
	}

	return topString, nil
}

// returnMainServerSection will return the string for the server
// section of the main.go file
func returnMainServerSection() (string, error) {
	topString := topCommentBlock +
		fmt.Sprintf("\n/	Server\n") +
		bottomCommentBlock +
		fmt.Sprintf("\n\nvar srv *http.Server\n\n") +
		fmt.Sprintf("// establishes the server contraints and information\n") +
		fmt.Sprintf("srv = &http.Server{\n") +
		fmt.Sprintf("Handler: handlers.CORS(handlers.AllowedHeaders([]string{\"origin\", \"X-Requested-With\", \"Content-Type\", \"Authorization\"}), handlers.AllowedMethods([]string{\"GET\", \"POST\", \"PUT\", \"HEAD\", \"OPTIONS\"}),\n") +
		fmt.Sprintf("handlers.AllowedOrigins([]string{\"*\"}),\n") +
		fmt.Sprintf("handlers.AllowCredentials())(r),\n") +
		fmt.Sprintf("Addr:         fmt.Sprintf(\":%%s\", port),\n") +
		fmt.Sprintf("WriteTimeout: 15 * time.Minute,\n") +
		fmt.Sprintf("ReadTimeout:  15 * time.Minute,\n}\n\n") +
		fmt.Sprintf("// runs the server\n") +
		fmt.Sprintf("log.Println(\"Running local on port: \", port)\n") +
		fmt.Sprintf("log.Fatal(srv.ListenAndServe())\n}")

	return topString, nil
}

// ReturnConfigFile will return the string for the config.go file
func ReturnConfigFile() ([]byte, error) {
	// structs
	structString, err := returnConfigStructsSection()
	if err != nil {
		// handle err
		log.Println(err)
	}
	// funcs
	funcString, err := returnConfigFuncsSection()
	if err != nil {
		// handle err
		log.Println(err)
	}

	return []byte(structString + funcString), nil
}

// returnConfigStructsSection will return the string for the structs
// found in the top of the config file
func returnConfigStructsSection() (string, error) {
	topString := fmt.Sprintf("package main\n\n") +
		fmt.Sprintf("import (\n") +
		fmt.Sprintf("\"fmt\"\n") +
		fmt.Sprintf("\"os\"\n") +
		fmt.Sprintf(")\n\n") +
		fmt.Sprintf("// Config is the struct to define our applciation configuration\n") +
		fmt.Sprintf("type Config struct {\n") +
		fmt.Sprintf("	Port     string         `json:\"port\"`\n") +
		fmt.Sprintf("	Env      string         `json:\"env\"`\n") +
		fmt.Sprintf("	Database PostgresConfig `json:\"database\"`\n}\n\n")

	postgreString := fmt.Sprintf("// PostgresConfig is the struct to define our postgres configuration\n") +
		fmt.Sprintf("type PostgresConfig struct {\n") +
		fmt.Sprintf("	Host     string `json:\"host\"`\n") +
		fmt.Sprintf("	Port     int    `json:\"port\"`\n") +
		fmt.Sprintf("	User     string `json:\"user\"`\n") +
		fmt.Sprintf("	Password string `json:\"password\"`\n") +
		fmt.Sprintf("	Dbname   string `json:\"dbname\"`\n}\n")

	return topString + postgreString, nil
}

// returnConfigFuncsSection will return the string for the funcs
// found in the top of the config file
func returnConfigFuncsSection() (string, error) {
	dialectString := fmt.Sprintf("\n// Dialect states that we are utilizing postrges\n") +
		fmt.Sprintf("func (c PostgresConfig) Dialect() string {\n") +
		fmt.Sprintf("	return \"postgres\"\n}\n\n")

	connectionString := fmt.Sprintf("\n// Connection makes the db connection string\n") +
		fmt.Sprintf("func (c PostgresConfig) Connection() string {\n") +
		fmt.Sprintf("	if c.Host == \"localhost\" {\n") +
		fmt.Sprintf("		if c.Password == \"\" {\n") +
		fmt.Sprintf("			return fmt.Sprintf(\"host=%%s port=%%d user=%%s dbname=%%s sslmode=disable\", c.Host, c.Port, c.User, c.Dbname)\n") +
		fmt.Sprintf("		}\n") +
		fmt.Sprintf("		return fmt.Sprintf(\"user=%%s password=%%s host=%%s dbname=%%s sslmode=disable\", c.User, c.Password, c.Host, c.Dbname)\n}\n") +
		fmt.Sprintf("	return fmt.Sprintf(\"user=%%s password=%%s host=%%s dbname=%%s\", c.User, c.Password, c.Host, c.Dbname)\n}\n\n")

	loadConfigString := fmt.Sprintf("// LoadConfig initializes the db based on env variables\n") +
		fmt.Sprintf("func LoadConfig() Config {\n") +
		fmt.Sprintf("	var c Config\n") +
		fmt.Sprintf("	var (\n") +
		fmt.Sprintf("		host = os.Getenv(\"DBHOST\")\n") +
		fmt.Sprintf("		name = os.Getenv(\"DBNAME\")\n") +
		fmt.Sprintf("		user = os.Getenv(\"DBUSER\")\n") +
		fmt.Sprintf("		pass = os.Getenv(\"DBPASS\")\n") +
		fmt.Sprintf("		port = os.Getenv(\"PORT\")\n") +
		fmt.Sprintf("		env  = os.Getenv(\"ENVIRONMENT\")\n	)\n\n") +
		fmt.Sprintf("	c.Database.Host = host\n") +
		fmt.Sprintf("	c.Database.Dbname = name\n") +
		fmt.Sprintf("	c.Database.User = user\n") +
		fmt.Sprintf("	c.Database.Password = pass\n") +
		fmt.Sprintf("	c.Database.Port = 5432\n\n") +
		fmt.Sprintf("	c.Port = port\n") +
		fmt.Sprintf("	c.Env = env\n\n") +
		fmt.Sprintf("	return c\n}")

	return dialectString +
		connectionString +
		loadConfigString, nil
}

// ReturnSecretStuffFile will return the string for the secretStuff.go file
func ReturnSecretStuffFile() ([]byte, error) {
	secretString, err := returnSetTheSecretStuff()
	if err != nil {
		// handle err
		log.Println(err)
	}

	return []byte(secretString), nil
}

// returnSetTheSecretStuff will return the string for the func
// that sets env variables
func returnSetTheSecretStuff() (string, error) {
	topString := fmt.Sprintf(`package main

	import (
		"os"
	)
	
	func setTheSecrets() {
		os.Setenv("DBHOST", "localhost")
		os.Setenv("DBNAME", "??")
		os.Setenv("DBUSER", "??")
		os.Setenv("DBPASS", "??")
		os.Setenv("PORT", "9001")
		os.Setenv("ENVIRONMENT", "??")
	}`)

	return topString, nil
}
