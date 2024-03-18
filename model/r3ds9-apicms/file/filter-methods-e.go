package file

//func (ca *Criteria) AndHexOIdEqTo(oId string) *Criteria {
//	const semLogContext = "mongo-file::and-hex-oid-eq-to"
//	if oId == "" {
//		return ca
//	}
//
//	objId, err := primitive.ObjectIDFromHex(oId)
//	if err != nil {
//		log.Error().Err(err).Msg(semLogContext)
//		return ca
//	}
//	mName := fmt.Sprintf(OID)
//	c := func() bson.E { return bson.E{Key: mName, Value: objId} }
//	*ca = append(*ca, c)
//	return ca
//}
