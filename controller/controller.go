package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

//DataValue as data
type DataValue struct{
	ID string
	Activity string
	Time string
}
type slicy []DataValue
//Database as dummy database
//var database map[string][]DataValue
var database = make(map[string]slicy)
var dataValue = DataValue{}


//Create will create a todo list
func Create(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method :",r.Method);
	if r.Method=="GET"{
		t,_ :=template.ParseFiles("template.html","formact.html","form.html")
		t.ExecuteTemplate(w,"form.html",nil)
	}else{
		r.ParseForm()
		if (r.FormValue("id") != "" &&  r.FormValue("activity") != "" &&  r.FormValue("time") != ""){
			dataValue.ID = r.FormValue("id")
			dataValue.Activity = r.FormValue("activity")	
			dataValue.Time = r.FormValue("time")	
		   
	   
			database["data"] = append(database["data"],dataValue)
	   
			   
			 t,_ := template.ParseFiles("form.html")
			 t.Execute(w,nil)
		}else{
			t,_ := template.ParseFiles("form.html")
			t.Execute(w,nil)
		}
	
	}
}


//ReadAll is an handler for reading all todo list
func ReadAll(w http.ResponseWriter, r *http.Request)  {
	t,_ :=template.ParseFiles("template.html","formact.html")
	t.Execute(w,database["data"])
}

//Update is used for updating a todo list
func Update(w http.ResponseWriter, r *http.Request)  {
	id:= mux.Vars(r)["id"]

	if r.Method == "GET" {
		for j, val := range database["data"]{
			if id == val.ID {
				t,_ := template.ParseFiles("formact.html");
				t.Execute(w,database["data"][j].ID);
			}
		};
	}else{
	activity := r.FormValue("activity");
	time := r.FormValue("time");
	for i, val := range database["data"]{
		if id == val.ID {
			database["data"][i].Activity = activity
			database["data"][i].Time = time
			t,_ := template.ParseFiles("template.html")
			t.Execute(w,database["data"])
		}
	}
	
	}

}

//Delete is used for deleting a todo list
func Delete(w http.ResponseWriter, r *http.Request)  {
	id:= mux.Vars(r)["id"]
	
	for i, val := range database["data"]{
		if id == val.ID {
			database["data"] = append(database["data"][:i],database["data"][i+1:]...)

		}
	}

	t,_ :=template.ParseFiles("template.html","formact.html","form.html")
	t.Execute(w,database["data"])
	

}
