package file

//const (
//	AddToSet UpdateOperator = "$addToSet"
//	Pull     UpdateOperator = "$pull"
//)
//
//func (ud *UpdateDocument) AddToSet() *Updates {
//	return ud.op(AddToSet)
//}
//
//func (ud *UpdateDocument) Pull() *Updates {
//	return ud.op(Pull)
//}
//
//func (ud *UpdateDocument) AddToEntRefsSet(p EntRefsStruct) *UpdateDocument {
//	ud.AddToSet().Add(func() bson.E {
//		return bson.E{Key: ENTREFS, Value: p}
//	})
//	return ud
//}
//
//func (ud *UpdateDocument) PullFromEntRefsSet(p EntRefsStruct) *UpdateDocument {
//	ud.Pull().Add(func() bson.E {
//		return bson.E{Key: ENTREFS, Value: p}
//	})
//	return ud
//}
