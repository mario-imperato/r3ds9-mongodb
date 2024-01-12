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
//----- objType - string -  [apps.[].objType apps.objType]

// AndAppsObjTypeEqTo No Remarks
func (ca *Criteria) AndAppsObjTypeEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(APPS_OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndAppsObjTypeIsNullOrUnset No Remarks
func (ca *Criteria) AndAppsObjTypeIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(APPS_OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAppsObjTypeIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(APPS_OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
