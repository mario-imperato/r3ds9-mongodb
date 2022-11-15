const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

print("domain collection - #docs: ", db.domain.countDocuments())
print("site collection - #docs: ", db.site.countDocuments())
print("session collection - #docs: ", db.session.countDocuments())
print("user collection - #docs: ", db.user.countDocuments())