package main


import (
	"net/http"
	"fmt"
	"strconv"
	"strings"
	"encoding/base64"
	"encoding/json"
)



//create hash handler
func HashCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		//error out
		fmt.Fprintf(w, "wrong request") // write data to response
	} else {
		r.ParseForm()
		// logic to create hash and jid
		//fmt.Println("password:", r.Form["password"])
		//
		var v = r.Form["password"]
		fmt.Println(v[0])

		var jid = jobRepo.NextId()

		//submit hashing job to worker pool/job channel
		job := HashingJob{jid, v[0]}

		go func() {
			fmt.Printf("added: %d\n", job.jid)
			jobChans <- job
		}()

		fmt.Fprintf(w, "jid: %d ", jid) // write data to response
	}

}

//show hash handler
func HashShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	if r.Method != "GET" {
		//error out
		fmt.Fprintf(w, "Invalid Request") // write data to response

	} else {

		var path = r.URL.Path
		var fields = strings.Split(path, "/")

		//todo better to use ohter Router/pattern match
		if len(fields) != 3 || len(fields[2]) == 0 {
			fmt.Fprintf(w, "Wrong Arguments") // write data to response

		} else {
			var jid = fields[2]
			fmt.Println("jid: " + jid)
			//look up the hash for jid
			var id int
			var err error
			if id, err = strconv.Atoi(jid); err != nil {
				//log error, set the correct htp status code
				fmt.Printf("invalid jid: %d \n", jid)
				panic(err)
			} else {
				var hashcode = jobRepo.Get(int32(id))
				if (hashcode == nil) {
					// If we didn't find it, 404

					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Not Found") // write data to response
				}
				fmt.Fprintf(w, base64.StdEncoding.EncodeToString(hashcode)) // write data to response
			}

		}
	}
}


func Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != "GET" {
		//error out
		//fmt.Fprintf(w, "wrong request") // write data to response
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotAcceptable, Text: "Invalid request"}); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Printf("stats: %d, %d \n", stats.Total, stats.Average)

		if err := json.NewEncoder(w).Encode(stats); err != nil {
			panic(err)
		}

	}
}

