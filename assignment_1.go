package main

import(
   "encoding/json"
   "net/http"
   "strings"
)
// Project struct for output
type Project struct {
	Repos           string `json:"repos"`
	Owner		string `json:"owner"`
	Committer	string `json:"committer"`
	Commits		int    `json:"commits"`
	Languages	[]string `json:"languages"`
}
// Contributors struct for the contributrs API
type Contributors struct {
	Login	        string `json:"login"`
	Contributions   int `json:"contributions"`
}
// Repos struct for the Repository API
type Repos struct {
        Owner	struct{
		   Login	 string `json:"login"`
		}
}

// GetContent was taken from the first answer on the question How to get JSON response in Golang: https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang

func getContent(url string, target interface{})error {

   res, err := http.Get(url)
   if err != nil{
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(target)
}

func handlerProjects(w http.ResponseWriter, r *http.Request){
   http.Header.Add(w.Header(), "content-type", "application/json")

   arg := r.URL.Path
   s := strings.Split(arg, "/")
   url := "https://api.github.com/repos/" + s[4] + "/" +s[5]
   langurl := url + "/languages"
   contrurl := url + "/contributors"

   ra := make(map[string]interface{})

   t1 := &Repos{}
   t2 := &[]Contributors{}

   t  := &Project{}

   err := getContent(url, t1)
   if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
   }

   err = getContent(contrurl, t2)
   if err != nil {
       http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
   }

   err = getContent(langurl, &ra)
      if err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
   }

   t.Repos = s[3] + "/" + s[4] + "/" + s[5]
   t.Owner = t1.Owner.Login
   t.Committer = (*t2)[0].Login
   t.Commits = (*t2)[0].Contributions

   lang := []string{}
    for i := range ra{
      lang = append(lang, i)
   }

   t.Languages = lang

   json.NewEncoder(w).Encode(t)
}

func main(){
   http.HandleFunc("/projectinfo/v1/", handlerProjects)
   http.ListenAndServe(":8080", nil)
}
