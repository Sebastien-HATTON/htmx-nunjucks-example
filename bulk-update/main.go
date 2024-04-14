package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	ID        int    `json:"id"           binding:"required"`
	FirstName string `json:"firstName"    `
	LastName  string `json:"lastName"     `
	Email     string `json:"email"        `
	Active    bool   `json:"active"       binding:"required"`
}

var contacts []Contact

func init_db() {

	contacts = []Contact{
		Contact{
			ID:        0,
			FirstName: "Joe",
			LastName:  "Smith",
			Email:     "joe@smith.org",
			Active:    true,
		},
		Contact{
			ID:        1,
			FirstName: "Angie",
			LastName:  "MacDowell",
			Email:     "angie@macdowell.org",
			Active:    true,
		},
		Contact{
			ID:        2,
			FirstName: "Fuqua",
			LastName:  "Tarkenton",
			Email:     "fuqua@tarkenton.org",
			Active:    true,
		},
		Contact{
			ID:        3,
			FirstName: "Kim",
			LastName:  "Yee",
			Email:     "kim@yee.org",
			Active:    false,
		},
	}
}

func main() {

	init_db()
	fmt.Println(contacts)

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, contacts)
	})

	router.POST("/users", func(c *gin.Context) {
		var updatedContacts []Contact

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		err = json.Unmarshal(body, &updatedContacts)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		for index, _ := range contacts {
			var active = false
			for _, newContact := range updatedContacts {
				if contacts[index].ID == newContact.ID {
					active = true
					break
				}
			}
			contacts[index].Active = active
		}
		c.JSON(http.StatusOK, contacts)
	})

	router.Run(":8080")
}
