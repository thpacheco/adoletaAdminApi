package authController

import (
	cosmoDb "adoletaAdminApi/db"
	userModel "adoletaAdminApi/models/users"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	jwtService "adoletaAdminApi/jwtSecurity"

	"go.mongodb.org/mongo-driver/bson"
)

var userCollection = cosmoDb.Db().Database("goTest").Collection("users")

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body userModel.User

	e := json.NewDecoder(r.Body).Decode(&body)

	if e != nil {
		fmt.Print(e)
	}

	err := userCollection.FindOne(context.TODO(), bson.D{{"username", body.UserName}, {"password", body.Password}}).Decode(&body)

	_e := json.NewDecoder(r.Body).Decode(&body)

	if _e != nil {

		var retToken = jwtService.GenerateJWT(body)
		json.NewEncoder(w).Encode(retToken)
	}

	if err != nil {
		fmt.Println(err)
	}

}
