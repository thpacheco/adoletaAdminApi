package authController

import (
	responseBase "adoletaAdminApi/common"
	cosmoDb "adoletaAdminApi/db"
	userModel "adoletaAdminApi/models/users"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

var userCollection = cosmoDb.Db().Database("adoleta").Collection("users")

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body userModel.User

	e := json.NewDecoder(r.Body).Decode(&body)

	if e != nil {
		fmt.Print(e)
	}

	err := userCollection.FindOne(context.TODO(), bson.D{{"username", body.UserName}, {"password", body.Password}}).Decode(&body)

	var response responseBase.Data

	if err != nil {
		fmt.Println(err)
		response = responseBase.Data{Success: false, Loading: true}
		json.NewEncoder(w).Encode(response)
	} else {
		response = responseBase.Data{Success: true, Loading: false}
		json.NewEncoder(w).Encode(response)

	}

}
