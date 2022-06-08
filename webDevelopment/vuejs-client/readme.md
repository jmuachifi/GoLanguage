# Creating your first VueJS client

- 1. Create a vuejs-client directory where we will keep all our VueJS source files and an HTTP server, as follows:
```
$ mkdir vuejs-client && cd vuejs-client && touch server.go
```
- 2. Copy the following code to server.go:
```
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"addEmployee",
		"POST",
		"/employee/add",
		addEmployee,
	},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
}
func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}
func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee  data :: ", err)
		return
	}
	log.Printf("adding employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id,
		employee.FirstName, employee.LastName)
	employees = append(employees, Employee{Id: employee.Id,
		FirstName: employee.FirstName, LastName: employee.LastName})
	json.NewEncoder(w).Encode(employees)
}
func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}

```
- 3. Create another directory with the name assets where all our frontend code files such as .html, .js, .css, and images will be kept, as follows:
```
$ mkdir assets && cd assets && touch index.html && touch main.js
```
- 4. Copy the following content to index.html:
```
<html>
  <head>
    <title>VueJs Client</title>
    <script type = "text/javascript" src = "https://cdnjs.
    cloudflare.com/ajax/libs/vue/2.4.0/vue.js"></script>
    <script type = "text/javascript" src="https://cdn.
    jsdelivr.net/npm/vue-resource@1.5.0"></script>
  </head>
  <body>
    <div id = "form">
      <h1>{{ message }}</h1>
      <table>
        <tr>
          <td><label for="id">Id</label></td>
          <td><input type="text" value="" v-model="id"/></td>
        </tr>
        <tr>
          <td><label for="firstName">FirstName</label></td>
          <td><input type="text" value="" v-model="firstName"/>
          <td>
        </tr>
        <tr>
          <td><label for="lastName">LastName</label></td>
          <td> <input type="text" value="" v-model="lastName" />
          </td>
        </tr>
        <tr>
          <td><a href="#" class="btn" @click="addEmployee">Add
          </a></td>
        </tr>
      </table>
    </div>
    <script type = "text/javascript" src = "main.js"></script>
  </body>
</html>
```
- 5. Copy the following content to main.js:
```
var vue_det = new Vue
({
 el: '#form',
 data: 
 {
   message: 'Employee Dashboard',
   id: '',
   firstName:'',
   lastName:''
 },
 methods: 
 {
   addEmployee: function() 
   {
     this.$http.post
     (
       '/employee/add', 
       {
         id: this.id,
         firstName:this.firstName,
         lastName:this.lastName
       }
     )
     .then
     (
       response => 
       {
         console.log(response);
       }, 
       error => 
       {
         console.error(error);
       }
     );
   }
 }
});
```
- 6. Run the program with the following command:
```
$ go run server.go
```