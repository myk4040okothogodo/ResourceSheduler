package data

import (
	"fmt"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/myk4040okothogodo/ResourceSheduler/timeslotsAPI/data"

)



// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrAssetNotFound = fmt.Errorf("Product not found")
// ErrAssetsNotFound is an error raised when  the db is empty
var ErrAssetsNotFound = fmt.Errorf("Assets not found")
//ErrAssetsNotCreated is an error raised when assets are not successfully created
var ErrAssetNotCreated = fmt.Errorf("Asset not created")
//ErrAssetsNotDeleted is an error raised when assets are not succesfully deleted
var ErrAssetNotDeleted = fmt.Errorf("Asset not Deleted")


type Asset struct {
  gorm.Model
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
  Owner User   `json:"owner" validate: "required"`
}


type User struct {
    gorm.Model
    name string `json: "name" gorm:"size:120;unique" validate: "required"`
    email string  `json: "email" gorm:"size:120"`
    phone_number string `json: "phone_number" gorm: "size: 20"`
}

func Init() (*gorm.DB,error) {
    dbURL := "host=localhost user=postgres password=mykokothe dbname=assetsgo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        return nil,err
    }

    db.AutoMigrate(&Asset{})
    return db, nil
}




type Connecting struct {
    DB *gorm.DB
}

func New(db *gorm.DB) Connecting {
    return Connecting{db}
}


// Product defines the structure for an API product
// swagger:model


// Products defines a slice of Product
type Assets []*Asset

// GetProducts returns all products from the database
func (c Connecting)GetAssets() (Assets, error) {
  
  var assets Assets
  if result := c.DB.Find(&assets); result.Error != nil {
      return nil,ErrAssetsNotFound
  }
	return assets, nil
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func (c Connecting)GetAssetByID(id int) (*Asset, error) {
	
  var asset *Asset
  if result := c.DB.First(&asset, id); result.Error != nil {
      return nil, ErrAssetNotFound
  }
  return asset, nil


  //i := findIndexByAssetID(id)
	//if id == -1 {
	//	return nil, ErrAssetNotFound
	//}

	//return assetList[i], nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func (c Connecting)UpdateAsset(p Asset, i int) error {

   if result := c.DB.First(&p, i); result.Error != nil {
       return  ErrAssetNotFound
   }
   
   c.DB.Save(&p)
   //log.println("Suceesfully updated")
	//i := findIndexByAssetID(p.ID)
	//if i == -1 {
	//	return ErrAssetNotFound
	//}

	// update the product in the DB
	//assetList[i] = &p

	return nil
}

// AddProduct adds a new product to the database
func (c Connecting)AddAsset(p Asset) error {

  if result := c.DB.Create(&p); result.Error != nil {
      return ErrAssetNotCreated
  }

  //log.println("Successfully added!")
	// get the next id in sequence
	//maxID := assetList[len(assetList)-1].ID
	//p.ID = maxID + 1
	//assetList = append(assetList, &p)
  return nil
}

// DeleteProduct deletes a product from the database
func (c Connecting)DeleteAsset(id int) error {

  var asset Asset;
  if result := c.DB.First(&asset, id); result.Error != nil {
      return ErrAssetNotDeleted;
  }
  //
  c.DB.Delete(&asset)

  //log.println("Successfully deleted!")
	//i := findIndexByAssetID(id)
	//if i == -1 {
	//	return ErrAssetNotFound
  return nil
}
	//assetList = append(assetList[:i], assetList[i+1])

	//return nil
//}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func (c Connecting)findIndexByAssetID(id int) int {
	for i, p := range assetList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var assetList = []*Asset{
	&Asset{
		ID:          1,
		Name:        "LandCruiser 200",
		Description: "A perfect off-roading machine",
		Price:        2400,
		SKU:         "abc323",
	},
	&Asset{
		ID:          2,
		Name:        "A CAT Grader",
		Description: "a road grading machine used for civil works",
		Price:       1999,
		SKU:         "fjd34",
	},
}


