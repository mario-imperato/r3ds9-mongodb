#!/bin/bash
mongosh mongodb://localhost:27017/r3ds9  --file site.js
mongosh mongodb://localhost:27017/r3ds9  --file domain.js
mongosh mongodb://localhost:27017/r3ds9  --file user.js
