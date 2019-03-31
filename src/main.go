package main

import(
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "cloud.google.com/go/bigtable"
  "cloud.google.com/go/storage"
)


// Ad features: will be used in the Post action to create an ad in the system
type Features struct{
  Id string `json:"id"`
  Description string `json:"description"`
  Url string `json:"url"`
  ImageUrl string `json:"image_url"`
}


// Ad interaction: save and read interactions for existing ads
// This is similar to Enum type
type Action string

const(
  CLICK_ACTION Action = "click"
  VIEW_ACTION Action = "view"
)

type InteractionRequest struct{
  Id string `json:"id"`
  Action Action `json:"action"`
}

type Interaction struct{
  View uint64 `json:"view"`
  Click uint64 `json:"action"`
}

// Ad structure: will be used in Get and List actions to return ad information back to client
type Ad struct{
  Features  Features  `json:"features"`
  Tags  map[string]float64  `json:"tags"`
  Interaction  Interaction  `json:"interaction"`
}


//
func main(){
  fmt.Println("started-service")

  r := mux.NewRouter()
  r.HandleFunc("/ad", postHandler).Methods("POST")
  r.HandleFunc("/ad", listHandler).Methods("GET")
  r.HandleFunc("/ad/{id}", getHandler).Methods("GET")
  r.HandleFunc("/ad/interaction", interactionHandler).Methods("POST")
  http.Handle("/", r)

  log.Fatal(http.ListenAndServe(":8080", nil))
}


// postHandler: handle post request from client, save a new add to its storage backend
func postHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("This is a post handler.")

  decoder := json.NewDecoder(r.Body)
  var f Features
  if err := decoder.Decode(&f); err != nil {
    http.Error(w, "Failed to parse JSON input", http.StatusBadRequest)
    fmt.Printf("Failed to parse JSON input %v\n", err)
    return
  } 
  fmt.Fprintf(w, "Post received: %s/n", f.Description)
}

// listHandler: handle list request from client and return all ads from storage backend back to client
func listHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("This is a list handler")

  // Return a list of fake ads
  var ads []*Ad
  ads = append(ads, &Ad{// codestyle: need a space here?
    Features: Features{
      Id: "1111",
      Description: "This is the first Ad",
      Url: "www.katrinayang.com",
      ImageUrl: "www.laioffer.com/images/1111",
    },
    Tags: map[string]float64{"foo": 1,"bar": 2},
    Interaction: Interaction{
      View: 2,
      Click: 1,
    },
  })

  ads = append(ads, &Ad{
    Features: Features{
      Id: "2222",
      Description: "this is the second Ad",
      Url: "www.laioffer.com",
      ImageUrl: "www.laioffer.com/images/2222",
    },
    Tags: map[string]float64{"foo": 1, "bar": 2},
    Interaction: Interaction{
      View: 1,
      Click: 0,
    },
  })

  js, err := json.Marshal(ads)
  if err != nil {
    m := fmt.Sprintf("Failed to parse ad objects to JSON %v", err)
    http.Error(w, m, http.StatusInternalServerError)
    return
  }

  w.Write(js)
}

// getHandler: handle get handler from client and return the corresponding ad back to client based on the Id in the request
func getHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("This is a get handler.")

  id := mux.Vars(r)["id"]
  fmt.Printf("Ad id is: %s", id)

  ad := &Ad{
    Features: Features{
      Id: id,
      Description: "this is the second Ad",
      Url: "www.laioffer.com",
      ImageUrl: "www.laioffer.com/image/2222",
    },
    Tags: map[string]float64{"foo": 1, "bar": 2},
    Interaction: Interaction{
      View: 1,
      Click: 0,
    },
  }

  js, err := json.Marshal(ad)
  if err != nil {
    m := fmt.Sprintf("Failed to parse Ad object to JSON %v", err)
    http.Error(w, m, http.StatusInternalServerError)
    return
  }

  w.Write(js)
}

// interactionHandler: 
func interactionHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("This is an interaction handler.")

  decoder := json.NewDecoder(r.Body)
  var req InteractionRequest
  err := decoder.Decode(&req)
  if err != nil {
    http.Error(w, "Failed to parse JSON input", http.StatusBadRequest)
    fmt.Printf("Failed to parse JSON input %v\n", err)
    return
  }
  fmt.Fprintf(w, "Interaction received: %s\n", req.Action)
}
