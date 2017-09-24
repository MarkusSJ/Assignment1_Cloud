package main

import(
   "testing"
   "net/http"
   "net/http/httptest"
)


func Test_handler(t *testing.T) {
   ts := httptest.NewServer(http.HandlerFunc(handlerProjects))
   defer ts.Close()

   res, err := http.Get(ts.URL + "/projectinfo/v1/github.com/apache/kafka")
   if err != nil {
      t.Errorf("Error making the GET request, %s", err)
   }
   if res.StatusCode != http.StatusOK {
      t.Errorf("Error, expected status OK got %d", res.StatusCode)

   }
}


