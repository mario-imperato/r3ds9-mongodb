#!/bin/bash
mongosh mongodb://localhost:27017/r3ds9  --file apigtw-domain.js
mongosh mongodb://localhost:27017/r3ds9  --file apigtw-site.js
mongosh mongodb://localhost:27017/r3ds9  --file apigtw-user.js
mongosh mongodb://localhost:27017/r3ds9  --file apigtw-session.js
mongosh mongodb://localhost:27017/r3ds9  --file apigtw-count-documents.js