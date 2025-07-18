package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/scraly/elephpants-api/pkg/swagger/server/models"
	"github.com/scraly/elephpants-api/pkg/swagger/server/restapi"

	"github.com/scraly/elephpants-api/pkg/swagger/server/restapi/operations"
	"github.com/scraly/elephpants-api/pkg/swagger/server/restapi/operations/elephpants"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewElephpantsAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	api.ElephpantsCheckHealthHandler = elephpants.CheckHealthHandlerFunc(Health)

	api.ElephpantsGetElephpantsHandler = elephpants.GetElephpantsHandlerFunc(GetElephpants)

	api.ElephpantsGetElephpantHandler = elephpants.GetElephpantHandlerFunc(GetElephpantByName)

	api.ElephpantsPostElephpantHandler = elephpants.PostElephpantHandlerFunc(CreateElephpant)

	api.ElephpantsDeleteElephpantHandler = elephpants.DeleteElephpantHandlerFunc(DeleteElephpant)

	api.ElephpantsPutElephpantHandler = elephpants.PutElephpantHandlerFunc(UpdateElephpant)

	api.ElephpantsGetElephpantImageHandler = elephpants.GetElephpantImageHandlerFunc(GetElephpantImageByName)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

type elephpant struct {
	Name        string `json:"name"`
	Displayname string `json:"displayname"`
	URL         string `json:"url"`
}

type allElephpants []elephpant

var theElephpants = allElephpants{
	{
		Name:        "ElePHPant",
		Displayname: "ElePHPant.png",
		URL:         "https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/ElePHPant.png",
	},
	{
		Name:        "ElePHPant-harry-potter",
		Displayname: "ElePHPant harry potter",
		URL:         "https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/ElePHPant-harry-potter.png",
	},
	{
		Name:        "forumphp",
		Displayname: "ForumPHP",
		URL:         "https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/forumphpv2.png",
	},
	{
		Name:        "yoda",
		Displayname: "Yoda PHP",
		URL:         "https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/yoda.png",
	},
	{
		Name:        "yodav2",
		Displayname: "Yoda PHP v2",
		URL:         "https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/yodav2.png",
	},
}

// Health route returns OK
func Health(elephpants.CheckHealthParams) middleware.Responder {
	fmt.Println("[Health] Call method")
	return elephpants.NewCheckHealthOK().
		WithPayload("OK").
		WithAccessControlAllowOrigin("*")
}

// Returns a a list of Elephpants
func GetElephpants(theElephpant elephpants.GetElephpantsParams) middleware.Responder {
	fmt.Println("[GetElephpants] Call method")

	var theElephpantsList []*models.Elephpant

	// Get all existing Elephpants
	for _, myElephpant := range theElephpants {
		theElephpantsList = append(theElephpantsList, &models.Elephpant{Name: myElephpant.Name, Displayname: myElephpant.Displayname, URL: myElephpant.URL})
	}

	return elephpants.NewGetElephpantsOK().
		WithPayload(theElephpantsList).
		WithAccessControlAllowOrigin("*")
}

// Returns an object of type Elephpant with a given name
func GetElephpantByName(elephpantParam elephpants.GetElephpantParams) middleware.Responder {
	fmt.Println("[GetElephpantByName] Call method")

	for _, myElephpant := range theElephpants {
		if myElephpant.Name == elephpantParam.Name {
			fmt.Println("Elephpant", elephpantParam.Name, "found")

			return elephpants.NewGetElephpantOK().WithPayload(
				&models.Elephpant{
					Name:        myElephpant.Name,
					Displayname: myElephpant.Displayname,
					URL:         myElephpant.URL}).
				WithAccessControlAllowOrigin("*")
		}
	}

	//If theElephpant have not been found, returns a 404 HTTP Error Code
	return elephpants.
		NewGetElephpantNotFound().
		WithAccessControlAllowOrigin("*")
}

// Returns an object of type Elephpant with a given name
func GetElephpantImageByName(elephpantParam elephpants.GetElephpantImageParams) middleware.Responder {
	fmt.Println("[GetElephpantByName] Call method")

	for _, myElephpant := range theElephpants {
		if myElephpant.Name == elephpantParam.Name {
			fmt.Println("Elephpant", elephpantParam.Name, "found")

			URL := "https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/" + elephpantParam.Name + ".png"

			// Get the image and return it
			response, err := http.Get(URL)
			if err != nil {
				fmt.Println(err)
			}
			defer response.Body.Close()

			if response.StatusCode == 200 {
				image, err := io.ReadAll(response.Body)
				if err != nil {
					fmt.Println(err)
				}

				return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
					rw.WriteHeader(200)
					rw.Write(image)
				})
			} else {
				fmt.Println("Error: " + elephpantParam.Name + " not exists! :-(")
			}
		}
	}

	//If theElephpant have not been found, returns a 404 HTTP Error Code
	return elephpants.
		NewGetElephpantNotFound().
		WithAccessControlAllowOrigin("*")
}

// TODO: to finish
func getElephpant(theElephpantName string) elephpant {
	for _, myElephpant := range theElephpants {
		if myElephpant.Name == theElephpantName {
			return myElephpant
		}
	}

	return elephpant{}
}

func theElephpantExists(theElephpantName string) bool {
	for _, myElephpant := range theElephpants {
		if myElephpant.Name == theElephpantName {
			return true
		}
	}

	return false
}

// Add a new Elephpant
func CreateElephpant(elephpantParam elephpants.PostElephpantParams) middleware.Responder {
	fmt.Println("[CreateElephpant] Call method")

	name := elephpantParam.Elephpant.Name
	displayname := elephpantParam.Elephpant.Displayname
	url := elephpantParam.Elephpant.URL

	fmt.Println("Try to create a Elephpant with the parameters:", *name, *displayname, *url)

	// Check if a theElephpant not already exists
	if !theElephpantExists(*name) {
		// Add new theElephpant in the list of existing Elephpants
		theElephpants = append(theElephpants, elephpant{*name, *displayname, *url})

		fmt.Println("Elephpant", *name, "created!")

		return elephpants.NewPostElephpantCreated().WithPayload(&models.Elephpant{Name: *name, Displayname: *displayname, URL: *url})
	} else {
		return elephpants.NewPostElephpantConflict()
	}
}

// Delete a Elephpant with a given name
func DeleteElephpant(elephpantParam elephpants.DeleteElephpantParams) middleware.Responder {
	fmt.Println("[DeleteElephpant] Call method")

	for i, myElephpant := range theElephpants {
		if myElephpant.Name == elephpantParam.Name {
			fmt.Println("Elephpant", elephpantParam.Name, "found, try to delete it")

			theElephpants = append(theElephpants[:i], theElephpants[i+1:]...)

			fmt.Println("Elephpant", elephpantParam.Name, "deleted!")

			return elephpants.NewDeleteElephpantOK()
		}
	}

	fmt.Println("[DeleteElephpant] End of the method")

	//If elephpant have not been found, returns a 404 HTTP Error Code
	return elephpants.NewDeleteElephpantNotFound()
}

// Update the displayname and the URL of an existing Elephpant
func UpdateElephpant(elephpantParam elephpants.PutElephpantParams) middleware.Responder {
	fmt.Println("[UpdateElephpant] Call method")

	fmt.Println("Updating", *elephpantParam.Elephpant.Name, "with new values")

	for i := range theElephpants {
		if theElephpants[i].Name == *elephpantParam.Elephpant.Name {
			theElephpants[i].Displayname = *elephpantParam.Elephpant.Displayname
			theElephpants[i].URL = *elephpantParam.Elephpant.URL

			fmt.Println("Elephpant updated!")

			return elephpants.NewPostElephpantCreated().WithPayload(&models.Elephpant{
				Name:        *elephpantParam.Elephpant.Name,
				Displayname: *elephpantParam.Elephpant.Displayname,
				URL:         *elephpantParam.Elephpant.URL})
		}
	}

	fmt.Println("[UpdateElephpant] End of the method")

	return elephpants.NewPutElephpantOK()
}

//TODO: Create Helper function in order to create a JSON with full existing Elephpants in github.com/scraly/theElephpants
// /*
// *
// Get Elephpants List from Scraly repository
// */
// func GetElephpantsList() []*models.Elephpant {

// 	client := github.NewClient(nil)

// 	// list public repositories for org "github"
// 	ctx := context.Background()
// 	// list all repositories for the authenticated user
// 	_, directoryContent, _, err := client.Repositories.GetContents(ctx, "scraly", "theElephpants", "/", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var arr []*models.Elephpant

// 	for _, c := range directoryContent {
// 		if *c.Name == ".gitignore" || *c.Name == "README.md" {
// 			continue
// 		}

// 		var name string = strings.Split(*c.Name, ".")[0]

// 		arr = append(arr, &models.Elephpant{name, *c.Displayname, *c.DownloadURL})

// 	}

// 	return arr
// }
