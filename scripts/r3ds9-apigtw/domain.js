const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "domain"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

let c = db.domain
if (!c)  {
    db.createCollection(r3ds9CollectionName)
}
else
{
    c.deleteMany({});
}

db[r3ds9CollectionName].insertOne(
    {
        "code" : "root",
        "objType": "domain",
        "name" : "Master Domain",
        "description" : "the master of domains",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-www"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "www-regatta"
            }
        ],
        "members": [
            {
               "code": "cvf", "objType": "domain"
            }
        ],
        "sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })

db[r3ds9CollectionName].insertOne(
    {
        "code" : "cvf",
        "objType": "domain",
        "name" : "Circolo Velico Fiumicino ASD",
        "description" : "Darsena Fiumicino",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-www"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "www-regatta"
            }
        ],
        "members": [
            {
                "code": "champ42", "objType": "site"
            }
        ],
        "sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })
