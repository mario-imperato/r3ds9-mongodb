package keyvaluepackage

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

//----- domain of type string
//----- domain - string -  [domain]

// AndDomainEqTo No Remarks
func (ca *Criteria) AndDomainEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(DOMAIN)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndDomainIsNullOrUnset No Remarks
func (ca *Criteria) AndDomainIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(DOMAIN)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndDomainIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(DOMAIN)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- site of type string
//----- site - string -  [site]

// AndSiteEqTo No Remarks
func (ca *Criteria) AndSiteEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(SITE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndSiteIsNullOrUnset No Remarks
func (ca *Criteria) AndSiteIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(SITE)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndSiteIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(SITE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
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

//----- category of type string
//----- category - string -  [category]

// AndCategoryEqTo No Remarks
func (ca *Criteria) AndCategoryEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(CATEGORY)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndCategoryIsNullOrUnset No Remarks
func (ca *Criteria) AndCategoryIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(CATEGORY)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndCategoryIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(CATEGORY)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
