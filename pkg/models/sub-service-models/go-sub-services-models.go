package subservices

import (
	"fmt"
	models "micro-gen/pkg/shared/models"
	"strings"
)

/**********************************************************************
/
/	Sub Service Section
/
/**********************************************************************/

// returnGoSubServiceTop will return the string of the top section of the
// SubService file
func returnGoSubServiceTop(name, projectName string) (string, error) {
	// main service fun
	topString := fmt.Sprintf("package services\n\n") +
		fmt.Sprintf("import (\n") +
		fmt.Sprintf("	\"%s/pkg/db\"\n", projectName) +
		fmt.Sprintf("	\"%s/pkg/validation\"\n)\n\n", projectName) +
		models.TopCommentBlock +
		fmt.Sprintf("\n/ Only Change this section if you are adding a new capability onto") +
		fmt.Sprintf("\n/ this subservice. Adding any new capability may nessesitate a change") +
		fmt.Sprintf("\n/ in the services.go With%s func found in pkg/sub-services/services.go", strings.Title(name)) +
		fmt.Sprintf("\n/ If you do not understand what is happening here, look into") +
		fmt.Sprintf("\n/ Interface Chaining. \n") +
		models.BottomCommentBlock
	return topString, nil
}

// returnGoSubServiceSvcDefinition will return the string of the service
// definition for the subService
func returnGoSubServiceSvcDefinition(name string) (string, error) {
	// service types
	srvcDefinitionString := fmt.Sprintf("\n\n // New%s loads related SQL statements and initializes the container struct\n", strings.Title(name)+"Service") +
		fmt.Sprintf("func New%s(s *Services) %s {\n", strings.Title(name)+"Service", strings.Title(name)+"Service") +
		fmt.Sprintf("	// create initial interface \n") +
		fmt.Sprintf("	ctx := &db.%s{}\n", strings.Title(name)+"StructDB") +
		fmt.Sprintf("	ctx.DB = s.db \n") +
		fmt.Sprintf("	srvc := &%s{}\n", strings.ToLower(name)+"Service") +
		fmt.Sprintf("	srvc.I%s = &validation.%s{I%s: ctx}\n", strings.Title(name)+"DB", strings.Title(name)+"Validator", strings.Title(name)+"DB") +
		fmt.Sprintf("	return srvc\n }\n\n") +
		// interface type
		fmt.Sprintf("// %s is a wrapper for related components\n", strings.Title(name)+"Services") +
		fmt.Sprintf("type %s interface {\n", strings.Title(name)+"Service") +
		fmt.Sprintf("	db.I%s\n}\n\n", strings.Title(name)+"DB") +
		// struct type
		fmt.Sprintf("type %s struct {\n", strings.ToLower(name)+"Service") +
		fmt.Sprintf("	db.I%s\n}\n", strings.Title(name)+"DB")

	return srvcDefinitionString, nil
}

/**********************************************************************
/
/	Service File Section
/
/**********************************************************************/

// returnServiceFileTopSection will return the string for the top
// section of the services.go file
func returnServiceFileTopSection() (string, error) {
	return "", nil
}

// returnServiceFileWithSection will return the string for the with
// funcs that declare the subservices in the services.go file
// and are later used in the /run/main.go file
func returnServiceFileWithSection(hasDB bool) (string, error) {
	topString := fmt.Sprintf("package services\n\n") +
		fmt.Sprintf("import (\n") +
		fmt.Sprintf("\"fmt\"\n") +
		fmt.Sprintf("\"log\"\n\n")
	if hasDB {
		topString += fmt.Sprintf("\"github.com/jmoiron/sqlx\"\n") +
			fmt.Sprintf("_ \"github.com/lib/pq\" // should be here\n\n") +
			fmt.Sprintf("migrate \"github.com/rubenv/sql-migrate\"\n")
	}
	topString += fmt.Sprintf(")")

	return topString, nil
}

// returnServiceFileNewServiceSection will return the string for the
// NewServices func that is found in the services.go file
func returnServiceFileNewServiceSection(hasDB bool, serviceArray []string) (string, error) {
	topString := fmt.Sprintf("\n\n// Config states how to define the with whatever\n") +
		fmt.Sprintf("// to be utilized in the NewServices func\n") +
		fmt.Sprintf("type Config func(*Services) error\n")

	withStrings := fmt.Sprintf("\n")
	if hasDB {
		// add postgres implementation
		topString += fmt.Sprintf("\n// WithPostgres returns a func to add a db onto Services\n") +
			fmt.Sprintf("func WithPostgres(dialect string, connectionInfo string) Config {\n") +
			fmt.Sprintf("	return func(s *Services) error {\n") +
			fmt.Sprintf("		db, err := sqlx.Open(dialect, connectionInfo)\n") +
			fmt.Sprintf("		if err != nil {\n") +
			fmt.Sprintf("			log.Println(\"Error connecting to DB\")\n") +
			fmt.Sprintf("			log.Println(err)\n") +
			fmt.Sprintf("			return err\n") +
			fmt.Sprintf("		}\n\n") +
			fmt.Sprintf("		s.db = db\n") +
			fmt.Sprintf("		return nil\n") +
			fmt.Sprintf("		}\n}")
	}
	// range on the serviceArray Strings to generate the withFuncs
	for _, v := range serviceArray {
		topString += fmt.Sprintf("\n// With%s returns a func to add the %s subservice onto Services\n", strings.Title(v), strings.ToLower(v)) +
			fmt.Sprintf("func With%s() Config {\n", strings.Title(v)) +
			fmt.Sprintf("	return func(s *Services) error {\n") +
			fmt.Sprintf("		s.%s = New%sService(s)\n", strings.Title(v), strings.Title(v)) +
			fmt.Sprintf("		return nil\n") +
			fmt.Sprintf("	}\n}")
	}
	commentBlock := fmt.Sprintf("\n") +
		models.TopCommentBlock +
		fmt.Sprintf("\n/ 	Shouldn't need to change this.\n") +
		models.BottomCommentBlock

	structString := fmt.Sprintf("\n\n// NewServices injects the service struct pointer to the above with funcs\n") +
		fmt.Sprintf("func NewServices(cfgs ...Config) (*Services, error) {\n") +
		fmt.Sprintf("	var s Services\n") +
		fmt.Sprintf("	for _, cfg := range cfgs {\n") +
		fmt.Sprintf("		if err := cfg(&s); err != nil {\n") +
		fmt.Sprintf("			return nil, err\n		}\n	}\n") +
		fmt.Sprintf("	return &s, nil\n") +
		fmt.Sprintf("}\n\n")

	return topString +
		withStrings +
		commentBlock +
		structString, nil
}

// returnServiceFileServicesStructSection will return the string for the
// services struct that is found in the services.go file
func returnServiceFileServicesStructSection(hasDB bool, serviceArray []string) (string, error) {
	commentBlock := models.TopCommentBlock +
		fmt.Sprintf("\n/	Add a field on the Services declaration when adding a new subService\n") +
		models.BottomCommentBlock

	servicesStruct := fmt.Sprintf("\n\n// Services is all the sub services within this service\n") +
		fmt.Sprintf("type Services struct {\n")

	for _, v := range serviceArray {
		servicesStruct += fmt.Sprintf("	%s	%sService\n", strings.Title(v), strings.Title(v))
	}

	if hasDB {
		servicesStruct += fmt.Sprintf("	db	*sqlx.DB\n}\n\n")
	}

	return commentBlock + servicesStruct, nil
}

// returnServiceFileGlobalSection will return the string for the global
// section that is found at the end of the services.go file and the methods
// are set onto the services struct
func returnServiceFileGlobalSection(hasDB bool) (string, error) {
	globalString := models.TopCommentBlock +
		fmt.Sprintf("\n/	Add any global methods that services will have access to\n") +
		models.BottomCommentBlock

	if hasDB {
		// Close
		globalString += fmt.Sprintf("\n\n// Close will close the database connection\n") +
			fmt.Sprintf("func (s *Services) Close() error {\n") +
			fmt.Sprintf("	return s.db.Close()\n}\n")

		//Migrate Up
		globalString += fmt.Sprintf("\n// MigrateDBUP runs the migration files up\n") +
			fmt.Sprintf("func (s *Services) MigrateDBUP() error {\n") +
			fmt.Sprintf("	// run the migrate here\n") +
			fmt.Sprintf("	migrations := &migrate.FileMigrationSource{\n") +
			fmt.Sprintf("		Dir: \"../resources/migrations\",\n	}\n") +
			fmt.Sprintf("	n, err := migrate.Exec(s.db.DB, \"postgres\", migrations, migrate.Up)\n") +
			fmt.Sprintf("	if err != nil {\n") +
			fmt.Sprintf("		// Handle errors!\n") +
			fmt.Sprintf("		log.Println(err)\n	}\n\n") +
			fmt.Sprintf("	fmt.Printf(\"Applied %%d migrations!\\n\", n)\n") +
			fmt.Sprintf("	return nil\n} \n")

		// Migrate Down
		globalString += fmt.Sprintf("\n//MigrateDBDown runs the migration files down\n") +
			fmt.Sprintf("func (s *Services) MigrateDBDown() error {\n") +
			fmt.Sprintf("	// run the migrate here\n") +
			fmt.Sprintf("	migrations := &migrate.FileMigrationSource{\n") +
			fmt.Sprintf("		Dir: \"../resources/migrations\",\n	}\n") +
			fmt.Sprintf("	n, err := migrate.Exec(s.db.DB, \"postgres\", migrations, migrate.Down)\n") +
			fmt.Sprintf("	if err != nil {\n") +
			fmt.Sprintf("		// Handle errors!\n") +
			fmt.Sprintf("		log.Println(err)\n}\n\n") +
			fmt.Sprintf("	fmt.Printf(\"Applied %%d migrations!\\n\", n)\n") +
			fmt.Sprintf("	return nil\n}\n")

		// helper to drop tables for faster local development
		globalString += fmt.Sprintf("\n// should remove, but helpfulish for clearing a db quickly locally\n") +
			fmt.Sprintf("/*\n") +
			fmt.Sprintf("DROP SCHEMA public CASCADE;\n") +
			fmt.Sprintf("CREATE SCHEMA public;\n") +
			fmt.Sprintf("GRANT ALL ON SCHEMA public TO postgres;\n") +
			fmt.Sprintf("GRANT ALL ON SCHEMA public TO public;\n") +
			fmt.Sprintf("*/\n")
	}
	return globalString, nil
}
