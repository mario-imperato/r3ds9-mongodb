package commons

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FilterGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

//----- objType of type string
//----- objType - string -  [app.objType]

// AndAppObjTypeEqTo No Remarks
func (ca *Criteria) AndAppObjTypeEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(APP_OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndAppObjTypeIsNullOrUnset No Remarks
func (ca *Criteria) AndAppObjTypeIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(APP_OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAppObjTypeIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(APP_OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
