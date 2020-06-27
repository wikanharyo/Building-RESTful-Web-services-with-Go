package main

import (
	"database/sql"
	"log"

	"encoding/json"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wikanharyo/Building-RESTful-Web-services-with-Go/dbutils"
)

// DB var holds global database driver
var DB *sql.DB

// Struct for database field
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

type StationResources struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

type ScheduleResources struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// Register adds paths and routes to restful container
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	// JSON only Content-Type
	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

// Handler functions
// GET http://localhost:9000/v1/trains/{train-id}
func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:9000/v1/trains
func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	statement, _ := DB.Prepare("INSERT INTO train  (DRIVER_NAME, OPERATING_STATUS) VALUES (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:9000/v1/trains/{train-id}
func (t TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare("DELETE FROM train WHERE id=?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	var err error
	// Connect to the database
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed")
	}
	// Create tables
	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("Start listening on localhost:9000")
	server := &http.Server{
		Addr:    ":9000",
		Handler: wsContainer,
	}
	log.Fatal(server.ListenAndServe())
}
