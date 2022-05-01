package handlers
import (
     "net/http"
     "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"
)

func (a *Assets) Update(rw http.ResponseWriter, r *http.Request){
   rw.Header().Add("Content-Type", "application/json")

   // Fetch the product from the context
   asst := r.Context().Value(KeyAsset{}).(data.Asset)
   a.l.Println("[DEBUG] updating record id", asst.ID)

   err := a.co.UpdateAsset(asst, asst.ID)
   if err == data.ErrAssetNotFound {
       a.l.Println("[ERROR] asset not found", err)
       rw.WriteHeader(http.StatusNotFound) 
       //data.ToJSON(&GenericError{message: "Product not found in database"}, rw)
       return
   }

   //Write the no content success header
   rw.WriteHeader(http.StatusNoContent)
}
