package file

//func FilterGoInfo() string {
//	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
//	return i
//}
//
////----- oId of type object-id
////----- oId - object-id -  [oId]
//
//// AndOIdEqTo No Remarks
//func (ca *Criteria) AndOIdEqTo(oId primitive.ObjectID) *Criteria {
//
//	if oId == primitive.NilObjectID {
//		return ca
//	}
//
//	mName := fmt.Sprintf(OID)
//	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
//	*ca = append(*ca, c)
//	return ca
//}
//
//func (ca *Criteria) AndOIdIn(p []primitive.ObjectID) *Criteria {
//
//	if len(p) == 0 {
//		return ca
//	}
//
//	mName := fmt.Sprintf(OID)
//	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
//	*ca = append(*ca, c)
//	return ca
//}
//
////----- ct of type string
////----- ct - string -  [vrnts.[].ct vrnts.ct]
//
//// AndVrntsCtEqTo No Remarks
//func (ca *Criteria) AndVrntsCtEqTo(p string) *Criteria {
//
//	if p == "" {
//		return ca
//	}
//
//	mName := fmt.Sprintf(VRNTS_CT)
//	c := func() bson.E { return bson.E{Key: mName, Value: p} }
//	*ca = append(*ca, c)
//	return ca
//}
//
//// AndVrntsCtIsNullOrUnset No Remarks
//func (ca *Criteria) AndVrntsCtIsNullOrUnset() *Criteria {
//
//	mName := fmt.Sprintf(VRNTS_CT)
//	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
//	*ca = append(*ca, c)
//	return ca
//}
//
//func (ca *Criteria) AndVrntsCtIn(p []string) *Criteria {
//
//	if len(p) == 0 {
//		return ca
//	}
//
//	mName := fmt.Sprintf(VRNTS_CT)
//	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
//	*ca = append(*ca, c)
//	return ca
//}
//
////----- cnt of type string
////----- cnt - string -  [vrnts.[].cnt vrnts.cnt]
//
//// AndVrntsCntEqTo No Remarks
//func (ca *Criteria) AndVrntsCntEqTo(p string) *Criteria {
//
//	if p == "" {
//		return ca
//	}
//
//	mName := fmt.Sprintf(VRNTS_CNT)
//	c := func() bson.E { return bson.E{Key: mName, Value: p} }
//	*ca = append(*ca, c)
//	return ca
//}
//
//// AndVrntsCntIsNullOrUnset No Remarks
//func (ca *Criteria) AndVrntsCntIsNullOrUnset() *Criteria {
//
//	mName := fmt.Sprintf(VRNTS_CNT)
//	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
//	*ca = append(*ca, c)
//	return ca
//}
//
//func (ca *Criteria) AndVrntsCntIn(p []string) *Criteria {
//
//	if len(p) == 0 {
//		return ca
//	}
//
//	mName := fmt.Sprintf(VRNTS_CNT)
//	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
//	*ca = append(*ca, c)
//	return ca
//}
