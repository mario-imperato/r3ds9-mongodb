### Schema file

#### Collection schema definition.

    {
	  "name": "file",
	  "properties": {
	    "folder-path": "./file",
	    "struct-name": "File"
	  },
	  "attributes": [
	    {
	      "name": "oId",
	      "type": "object-id",
	      "tags": [
	        "json",
	        "_id",
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
	      "name": "fn",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "descr",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "role",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "entRefs",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "struct-name": "EntRefsStruct",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "dom",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "ns",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "entType",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "entId",
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
	      "name": "metadata",
	      "type": "document",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "vrnts",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "struct-name": "Variant",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "ct",
	            "type": "string",
	            "queryable": true,
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "wd",
	            "type": "int",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "ht",
	            "type": "int",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "lks",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "bln",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "cnt",
	            "type": "string",
	            "queryable": true,
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "url",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "role",
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
	      "name": "sysinfo",
	      "type": "ref-struct",
	      "struct-ref": {
	        "struct-name": "SysInfo",
	        "is-external": true,
	        "Package": "github.com/mario-imperato/r3ds9-mongodb/model/commons",
	        "Item": null
	      }
	    }
	  ]
	}

