const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

let c = db.site
if (!c)  {
    db.createCollection('site')
}
else
{
    c.deleteMany({});
}

db.site.insertOne(
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
                ,"path": "www-regatta"
            }
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })
