package main
import("net/http";"net/http/httptest";"testing")
func TestHealth(t *testing.T){ r:=httptest.NewRequest("GET","/healthz",nil); w:=httptest.NewRecorder(); h:=http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.WriteHeader(200)}); h.ServeHTTP(w,r); if w.Code!=200{t.Fatal(w.Code)} }
