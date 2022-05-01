package handlers

import "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"



type errorResponseWrapper struct {
    // Description
    // in: body
    Body GenericError

}

//valiation errors defined as an array of strings
type errorValidationWrapper struct {
    Body ValidationError

}

type assetsResponsewrapper struct {
    // All current products
    // in : body
    Body []data.Asset
}

type assetResponseWrapper struct {
    // Newly created asset
    // in : body
    Body data.Asset

}

type noContentResponseWrapper struct {
}

type productIDParamsWrapper  struct {
    ID int `json:"id"`

}

