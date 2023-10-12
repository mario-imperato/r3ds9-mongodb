const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "apigtw_site"
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

db[r3ds9CollectionName].insertOne(
    {
        "code" : "champ42",
        "domain": "cvf",
        "objType": "site",
        "name" : "CVF Campionato 42",
        "description" : "Il campionato piu' affollato d'Italia",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-www"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "r3ds9-app-home-console/index.tmpl"
            }
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })
