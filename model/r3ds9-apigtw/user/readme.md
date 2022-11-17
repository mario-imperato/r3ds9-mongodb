### Schema user

#### Collection schema definition.

    {
	  "name": "user",
	  "properties": {
	    "folder-path": "./user",
	    "struct-name": "User"
	  },
	  "attributes": [
	    {
	      "name": "oId",
	      "type": "object-id",
	      "tags": [
	        "json",
	        "-",
	        "bson",
	        "_id"
	      ],
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "nickname",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "objType",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "firstname",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "lastnamw",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "password",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "roles",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "type": "ref-struct",
	        "struct-ref": {
	          "struct-name": "UserRole",
	          "is-external": true,
	          "Package": "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-core/commons",
	          "Item": null
	        }
	      },
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "sysinfo",
	      "type": "ref-struct",
	      "struct-ref": {
	        "struct-name": "SysInfo",
	        "is-external": true,
	        "Package": "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-core/commons",
	        "Item": null
	      }
	    }
	  ]
	}

