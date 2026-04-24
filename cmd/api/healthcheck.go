package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// v1
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)

	//v2
	// js := `{"status": "available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)

	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(js))

	//v3
	// data := map[string]string{
	// 	"status":      "available",
	// 	"environment": app.config.env,
	// 	"version":     version,
	// }

	// js, err := json.Marshal(data)
	// if err != nil {
	// 	app.logger.Error(err.Error())
	// 	http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	// 	return
	// }
	// js = append(js, '\n')
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
