const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

print("[apigtw] - domain collection - #docs: ", db.apigtw_domain.countDocuments())
print("[apigtw] - site collection - #docs: ", db.apigtw_site.countDocuments())
print("[apigtw] - session collection - #docs: ", db.apigtw_session.countDocuments())
print("[apigtw] - user collection - #docs: ", db.apigtw_user.countDocuments())
