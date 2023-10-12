package user

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

//----- oId of type object-id
//----- oId - object-id -  [oId]

// AndOIdEqTo No Remarks
func (ca *Criteria) AndOIdEqTo(oId primitive.ObjectID) *Criteria {

	if oId == primitive.NilObjectID {
		return ca
	}

	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndOIdIn(p []primitive.ObjectID) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- nickname of type string
//----- nickname - string -  [nickname]

// AndNicknameEqTo No Remarks
func (ca *Criteria) AndNicknameEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(NICKNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndNicknameIsNullOrUnset No Remarks
func (ca *Criteria) AndNicknameIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(NICKNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndNicknameIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(NICKNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndEmailEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(EMAIL)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

//----- objType of type string
//----- objType - string -  [objType]

// AndObjTypeEqTo No Remarks
func (ca *Criteria) AndObjTypeEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndObjTypeIsNullOrUnset No Remarks
func (ca *Criteria) AndObjTypeIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndObjTypeIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
