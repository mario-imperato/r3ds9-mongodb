const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

print("[apicms] user collection - #docs: ", db.apicms_user.countDocuments())

