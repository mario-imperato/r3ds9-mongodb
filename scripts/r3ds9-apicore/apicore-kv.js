const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "apicore_kv"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

let c = db[r3ds9CollectionName]
if (!c)  {
    db.createCollection(r3ds9CollectionName)
}
else
{
    c.deleteMany({});
}

c.insertOne(
    {
        "domain": "root",
        "site" : "*",
        "objType": "kv",
        "category": "console-ui-look",
        "issystem": true,
        "description": "Aspetto UI di amministrazione",
        "inherit": false
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

