// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "zang",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin": {
            "get": {
                "description": "列出所有管理员",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ListAdmin"
                ],
                "summary": "列出所有管理员",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "pageNo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "管理员ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "管理员名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "管理员邮箱",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "创建时间",
                        "name": "createdAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "更新时间",
                        "name": "updatedAt",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取所有管理员",
                        "schema": {
                            "$ref": "#/definitions/response.PageResp"
                        }
                    }
                }
            },
            "put": {
                "description": "编辑管理员",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EditAdmin"
                ],
                "summary": "编辑管理员",
                "parameters": [
                    {
                        "description": "Admin 信息",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AdminEditReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功编辑管理员",
                        "schema": {
                            "$ref": "#/definitions/model.AdminEditReq"
                        }
                    }
                }
            },
            "post": {
                "description": "添加新管理员",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AddAdmin"
                ],
                "summary": "添加新管理员",
                "parameters": [
                    {
                        "description": "Admin 信息",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AdminAddReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "成功添加管理员",
                        "schema": {
                            "$ref": "#/definitions/model.AdminAddReq"
                        }
                    }
                }
            }
        },
        "/admin/:id": {
            "delete": {
                "description": "删除指定id的管理员",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "DeleteAdmin"
                ],
                "summary": "删除指定id的管理员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功删除管理员",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "健康检查",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "健康检查",
                "responses": {
                    "200": {
                        "description": "健康检查成功",
                        "schema": {
                            "$ref": "#/definitions/model.HealthResp"
                        }
                    }
                }
            }
        },
        "/permission": {
            "get": {
                "description": "展示数据库中权限表中的所有权限",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ListPermission"
                ],
                "summary": "展示所有权限",
                "responses": {
                    "200": {
                        "description": "成功获取权限列表",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.PermissionsListReq"
                            }
                        }
                    },
                    "500": {
                        "description": "内部服务器错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role": {
            "get": {
                "description": "列出所有角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ListRole"
                ],
                "summary": "列出所有角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "pageNo",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色描述",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "创建时间",
                        "name": "createdAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "更新时间",
                        "name": "updatedAt",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取所有角色",
                        "schema": {
                            "$ref": "#/definitions/response.PageResp"
                        }
                    }
                }
            },
            "put": {
                "description": "编辑角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EditRole"
                ],
                "summary": "编辑角色",
                "parameters": [
                    {
                        "description": "Role 信息",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RolesEditReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功编辑角色",
                        "schema": {
                            "$ref": "#/definitions/model.RolesEditReq"
                        }
                    }
                }
            },
            "post": {
                "description": "添加新角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AddRole"
                ],
                "summary": "添加新角色",
                "parameters": [
                    {
                        "description": "Role 信息",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RolesAddReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "成功添加角色",
                        "schema": {
                            "$ref": "#/definitions/model.RolesAddReq"
                        }
                    }
                }
            }
        },
        "/role/:id": {
            "delete": {
                "description": "删除指定id的角色",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "DeleteRole"
                ],
                "summary": "删除指定id的角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功删除角色",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/site": {
            "get": {
                "description": "从数据库中获取所有站点记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Site"
                ],
                "summary": "获取所有站点",
                "responses": {
                    "200": {
                        "description": "成功获取所有站点",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.SitesResp"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "编辑site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EditSite"
                ],
                "summary": "编辑site",
                "parameters": [
                    {
                        "description": "Site 信息",
                        "name": "site",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SitesEditReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功编辑site",
                        "schema": {
                            "$ref": "#/definitions/model.SitesEditReq"
                        }
                    }
                }
            },
            "post": {
                "description": "添加新site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AddSite"
                ],
                "summary": "添加新site",
                "parameters": [
                    {
                        "description": "Site 信息",
                        "name": "site",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SitesAddReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "成功添加site",
                        "schema": {
                            "$ref": "#/definitions/model.SitesAddReq"
                        }
                    }
                }
            }
        },
        "/site/:id": {
            "get": {
                "description": "根据id获取site",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetSiteById"
                ],
                "summary": "根据id获取site",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取site",
                        "schema": {
                            "$ref": "#/definitions/model.SitesResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除指定id的site",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "DeleteSite"
                ],
                "summary": "删除指定id的site。",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功删除site",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tag": {
            "get": {
                "description": "获取数据库中所有标签记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "获取所有标签",
                "responses": {
                    "200": {
                        "description": "成功获取所有标签",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TagsResp"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "编辑tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EditTag"
                ],
                "summary": "编辑tag",
                "parameters": [
                    {
                        "description": "Tag 信息",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TagsEditReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功编辑tag",
                        "schema": {
                            "$ref": "#/definitions/model.TagsEditReq"
                        }
                    }
                }
            },
            "post": {
                "description": "添加新tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AddTag"
                ],
                "summary": "添加新tag",
                "parameters": [
                    {
                        "description": "Tag 信息",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TagsAddReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "成功添加tag",
                        "schema": {
                            "$ref": "#/definitions/model.TagsAddReq"
                        }
                    }
                }
            }
        },
        "/tag/:id": {
            "get": {
                "description": "根据id获取tag",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetTagById"
                ],
                "summary": "根据id获取tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取tag",
                        "schema": {
                            "$ref": "#/definitions/model.TagsResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除指定id的tag",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "DeleteTag"
                ],
                "summary": "删除指定id的tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功删除tag",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AdminAddReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.AdminEditReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "role_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.HealthResp": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "model.PermissionsListReq": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.RolesAddReq": {
            "type": "object"
        },
        "model.RolesEditReq": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.SitesAddReq": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.SitesEditReq": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.SitesResp": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.TagsAddReq": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.TagsEditReq": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.TagsResp": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "response.PageResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {},
                "pageNo": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8088",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
