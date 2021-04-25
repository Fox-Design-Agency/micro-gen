package routeHandlers

import (
	"fmt"
	models "micro-gen/pkg/shared/models"
	"strings"
)

// returnGoRouteHandlerTop will return the top section for the
// route handlers layer
func returnGoRouteHandlerTop(name, projectName string) (string, error) {
	topString := fmt.Sprintf("package routeHandlers\n\n") +
		fmt.Sprintf("import (\n") +
		fmt.Sprintf("	\"%s/pkg/helpers\"\n", projectName) +
		fmt.Sprintf("	services \"%s/pkg/sub-services\"\n", projectName) +
		fmt.Sprintf("	\"net/http\"\n)\n\n") +
		models.TopCommentBlock +
		fmt.Sprintf("\n/ Only Change this section if you are adding a new subservice onto") +
		fmt.Sprintf("\n/ this service. Adding any new subservice will nessesitate a change") +
		fmt.Sprintf("\n/ in the main.go route-handler declaration found in run/main.go\n") +
		models.BottomCommentBlock

	newStructDeclaration := fmt.Sprintf("\n\n// %s is a func that created a new %s struct.", "New"+strings.Title(name), strings.Title(name)) +
		fmt.Sprintf("\n// The %s struct requires:", strings.Title(name)) +
		fmt.Sprintf("\n// %s\n", strings.Title(name)+"Service") +
		fmt.Sprintf("func New%s(\n", strings.Title(name)) +
		fmt.Sprintf("	ss services.%s) *%s {\n", strings.Title(name)+"Service", strings.Title(name)) +
		fmt.Sprintf("	return &%s{\n", strings.Title(name)) +
		fmt.Sprintf("		ss: ss,\n") +
		fmt.Sprintf("	}\n}\n\n")

	structDeclaration := fmt.Sprintf("// %s defines our %s route-handler and\n", strings.Title(name), strings.ToLower(name)) +
		fmt.Sprintf("// the subservices that it has access to \n") +
		fmt.Sprintf("type %s struct {\n", strings.Title(name)) +
		fmt.Sprintf("	ss services.%s\n", strings.Title(name)+"Service") +
		fmt.Sprintf("}\n\n")

	return topString + newStructDeclaration + structDeclaration, nil
}

// returnGoRouteHandlerMethods will return the string for the method
// section definitions
func returnGoRouteHandlerMethods(name string, hasCRUD bool) (string, error) {

	methodBlock := models.TopCommentBlock +
		fmt.Sprintf("\n/	Add new methods onto %s that can be accessed through the %s\n", strings.Title(name), strings.ToLower(name)) +
		fmt.Sprintf("/ 	route-handler.\n") +
		models.BottomCommentBlock

	if hasCRUD {
		// Create
		methodBlock += fmt.Sprintf("\n\n// Create%s will create a %s\n", strings.Title(name), strings.ToLower(name)) +
			fmt.Sprintf("func (rh *%s) Create%s(w http.ResponseWriter, r *http.Request) {\n", strings.Title(name), strings.Title(name)) +
			fmt.Sprintf("// validate context \n") +
			fmt.Sprintf("//\n\n") +
			fmt.Sprintf("// Parse JSON \n") +
			fmt.Sprintf("//\n\n") +
			fmt.Sprintf("// Call DB layer \n") +
			fmt.Sprintf("if _, err := rh.ss.Create%s(); err != nil {\n", strings.Title(name)) +
			fmt.Sprintf("// Handle Error \n") +
			fmt.Sprintf("helpers.SendErrorHeader(w, 500, nil)\n") +
			fmt.Sprintf("return\n") +
			fmt.Sprintf("}\n\n") +
			fmt.Sprintf("// succeeded\n") +
			fmt.Sprintf("// should send success header\n") +
			fmt.Sprintf("helpers.SendSuccessHeader(w, nil) \n") +
			fmt.Sprintf("}\n\n")

		// Delete
		methodBlock += fmt.Sprintf("\n\n// Delete%s will delete a %s\n", strings.Title(name), strings.ToLower(name)) +
			fmt.Sprintf("func (rh *%s) Delete%s(w http.ResponseWriter, r *http.Request) {\n", strings.Title(name), strings.Title(name)) +
			fmt.Sprintf("// validate context \n") +
			fmt.Sprintf("//\n\n") +
			fmt.Sprintf("// get the query id \n") +
			fmt.Sprintf("id := r.FormValue(\"id\")\n\n") +
			fmt.Sprintf("// Call DB layer \n") +
			fmt.Sprintf("if _, err := rh.ss.Delete%s(id); err != nil {\n", strings.Title(name)) +
			fmt.Sprintf("// Handle Error \n") +
			fmt.Sprintf("helpers.SendErrorHeader(w, 500, nil)\n") +
			fmt.Sprintf("return\n") +
			fmt.Sprintf("}\n\n") +
			fmt.Sprintf("// succeeded\n") +
			fmt.Sprintf("// should send success header\n") +
			fmt.Sprintf("helpers.SendSuccessHeader(w, nil) \n") +
			fmt.Sprintf("}\n\n")

		// Get All
		methodBlock += fmt.Sprintf("\n\n// GetAll%s will get all %s\n", strings.Title(name), strings.ToLower(name)) +
			fmt.Sprintf("func (rh *%s) GetAll%s(w http.ResponseWriter, r *http.Request) {\n", strings.Title(name), strings.Title(name)) +
			fmt.Sprintf("// validate context \n") +
			fmt.Sprintf("//\n\n") +
			fmt.Sprintf("// get the query id \n") +
			fmt.Sprintf("id := r.FormValue(\"id\")\n\n") +
			fmt.Sprintf("// Call DB layer \n") +
			fmt.Sprintf("if _, err := rh.ss.GetAll%s(id); err != nil {\n", strings.Title(name)) +
			fmt.Sprintf("// Handle Error \n") +
			fmt.Sprintf("helpers.SendErrorHeader(w, 500, nil)\n") +
			fmt.Sprintf("return\n") +
			fmt.Sprintf("}\n\n") +
			fmt.Sprintf("// succeeded\n") +
			fmt.Sprintf("// should send success header\n") +
			fmt.Sprintf("helpers.SendSuccessHeader(w, nil) \n") +
			fmt.Sprintf("}\n\n")

		// Get
		methodBlock += fmt.Sprintf("\n\n// Get%s will get a single %s\n", strings.Title(name), strings.ToLower(name)) +
			fmt.Sprintf("func (rh *%s) Get%s(w http.ResponseWriter, r *http.Request) {\n", strings.Title(name), strings.Title(name)) +
			fmt.Sprintf("// validate context \n") +
			fmt.Sprintf("//\n\n") +
			fmt.Sprintf("// get the query id \n") +
			fmt.Sprintf("id := r.FormValue(\"id\")\n\n") +
			fmt.Sprintf("// Call DB layer \n") +
			fmt.Sprintf("if _, err := rh.ss.Get%s(id); err != nil {\n", strings.Title(name)) +
			fmt.Sprintf("// Handle Error \n") +
			fmt.Sprintf("helpers.SendErrorHeader(w, 500, nil)\n") +
			fmt.Sprintf("return\n") +
			fmt.Sprintf("}\n\n") +
			fmt.Sprintf("// succeeded\n") +
			fmt.Sprintf("// should send success header\n") +
			fmt.Sprintf("helpers.SendSuccessHeader(w, nil) \n") +
			fmt.Sprintf("}\n\n")

		// Update
		methodBlock += fmt.Sprintf("\n\n// Update%s will update %s\n", strings.Title(name), strings.ToLower(name)) +
			fmt.Sprintf("func (rh *%s) Update%s(w http.ResponseWriter, r *http.Request) {\n", strings.Title(name), strings.Title(name)) +
			fmt.Sprintf("// validate context \n") +
			fmt.Sprintf("//\n\n") +
			fmt.Sprintf("// get the query id \n") +
			fmt.Sprintf("id := r.FormValue(\"id\")\n\n") +
			fmt.Sprintf("// Parse JSON \n") +
			fmt.Sprintf("//\n\n") +
			fmt.Sprintf("// Call DB layer \n") +
			fmt.Sprintf("if _, err := rh.ss.Update%s(id); err != nil {\n", strings.Title(name)) +
			fmt.Sprintf("// Handle Error \n") +
			fmt.Sprintf("helpers.SendErrorHeader(w, 500, nil)\n") +
			fmt.Sprintf("return\n") +
			fmt.Sprintf("}\n\n") +
			fmt.Sprintf("// succeeded\n") +
			fmt.Sprintf("// should send success header\n") +
			fmt.Sprintf("helpers.SendSuccessHeader(w, nil) \n") +
			fmt.Sprintf("}\n\n")
	}

	return methodBlock, nil
}
