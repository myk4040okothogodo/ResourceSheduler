package handlers
import (
    "fmt"
    "context"
    "net/http"
    protos "github.com/myk4040okothogodo/ResourceSheduler/protos/assetsdefinition/protos"
    "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"
  
)


//listall handles GET requests and returns all current assets

func (a *Assets) ListAll(rw http.ResponseWriter, r *http.Request) {
    a.l.Println("[DEBUG] get all records")
    rw.Header().Add("Content-Type", "application/json")

    assts, _ := a.co.GetAssets()

    err := data.ToJSON(assts, rw)
    if err != nil {
        //we should never be here but log the error just incase
        a.l.Println("[ERROR] serializing asset", err)

    }
}

//ListsSingle handle GET request
func (a *Assets) ListSingle (rw http.ResponseWriter, r *http.Request){
    rw.Header().Add("Content-Type", "application/json")
    id := getAssetID(r)

    a.l.Println("[DEBUG] get record id", id)
    asst,err := a.co.GetAssetByID(id)

    switch err {
    case nil:

    case data.ErrAssetNotFound:
      a.l.Println("[ERROR] fetching assets", err)

      rw.WriteHeader(http.StatusNotFound)
      data.ToJSON(&GenericError{Message: err.Error()}, rw)
      return
    default:
      a.l.Println("[ERROR] fetching product", err)
      rw.WriteHeader(http.StatusInternalServerError)
      data.ToJSON(&GenericError{Message: err.Error()}, rw)
      return
    }

    // get asset schedules
    ar := &protos.GetAssetAvailabilityRequest{
      Assetcategory: protos.AssetCategory(protos.AssetCategory_value["Vehicular"]),
    }

    resp, err := a.ac.CheckForAssetAvailability(context.Background(), ar)
    if err != nil {
        a.l.Println("[error] error getting new rate", err)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    }

    a.l.Println("Resp %#v", resp)
    fmt.Println(resp)    
    err = data.ToJSON(asst, rw)
    if err != nil {
        // we should never be here but log the error just incase
        a.l.Println("[ERROR] serializing asset", err)
    }
}
