### Schema domain

#### Collection schema definition.

    {
	  "name": "domain",
	  "properties": {
	    "folder-path": "./domain",
	    "struct-name": "Domain"
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
	      "name": "code",
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
	      "name": "name",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "description",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "langs",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "members",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "struct-name": "DomainMember",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "code",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "objType",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          }
	        ],
	        "struct-ref": {
	          "Package": "",
	          "Item": null
	        }
	      },
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "apps",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "type": "ref-struct",
	        "struct-ref": {
	          "struct-name": "App",
	          "is-external": true,
	          "Package": "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apicore/commons",
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
	        "Package": "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apicore/commons",
	        "Item": null
	      }
	    }
	  ]
	}

