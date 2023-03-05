const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

print("domain collection - #docs: ", db.apicore_domain.countDocuments())
print("site collection - #docs: ", db.apicore_site.countDocuments())
print("session collection - #docs: ", db.apicore_session.countDocuments())
print("user collection - #docs: ", db.apicore_user.countDocuments())