const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

let c = db.domain
if (!c)  {
    db.createCollection('domain')
}
else
{
    c.deleteMany({});
}

db.domain.insertOne(
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
