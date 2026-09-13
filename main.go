package main

import (
 "database/sql"
 "encoding/json"
 "log"
 "net/http"
 _ "modernc.org/sqlite"
)

func main() {
 db, err := sql.Open("sqlite", "file:/data/race.db?cache=shared")
 if err != nil { log.Fatal(err) }
 defer db.Close()
 if _, err = db.Exec("create table if not exists audit_events(id integer primary key, payload text not null)"); err != nil { log.Fatal(err) }
 mux := http.NewServeMux()
 mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) { json.NewEncoder(w).Encode(map[string]string{"status":"ok"}) })
 mux.HandleFunc("POST /events", func(w http.ResponseWriter, r *http.Request) { var v map[string]any; if json.NewDecoder(r.Body).Decode(&v) != nil { http.Error(w, "invalid json", 400); return }; b,_:=json.Marshal(v); if _,e:=db.Exec("insert into audit_events(payload) values(?)",b); e!=nil { http.Error(w,"storage error",500); return }; w.WriteHeader(http.StatusCreated) })
 log.Fatal(http.ListenAndServe(":8080", mux))
}
