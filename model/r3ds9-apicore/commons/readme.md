### Schema commons

#### Collection schema definition.

    {
	  "name": "commons",
	  "properties": {
	    "folder-path": "./commons",
	    "struct-name": "Commons"
	  },
	  "attributes": [
	    {
	      "name": "sysinfo",
	      "struct-name": "SysInfo",
	      "type": "struct",
	      "attributes": [
	        {
	          "name": "status",
	          "type": "string",
	          "struct-ref": {
	            "Package": "",
	            "Item": null
	          }
	        },
	        {
	          "name": "createdat",
	          "type": "date",
	          "struct-ref": {
	            "Package": "",
	            "Item": null
	          }
	        },
	        {
	          "name": "modifiedat",
	          "type": "date",
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
	    {
	      "name": "apps",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "struct-name": "App",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "id",
	            "type": "string",
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
	            "name": "path",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "roleRequired",
	            "type": "bool",
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
	      "name": "roles",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "struct-name": "UserRole",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "domain",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "site",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "apps",
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
	    }
	  ]
	}

