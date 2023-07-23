# Headless CMS service

## Description

This is used for my personal portfolio website. It is essentially used, to create dynamic content for the website. This includes dynamic schema, dynamic content and dynamic endpoint creation.

### CMS

This is implemented using Rust and MongoDB.

#### Dynamic schema

1. create a new schema
2. add dynamic fields to the schema
3. save the schema

Dynamic data can be handled by using ``serde_json::Value`` or creating new structs:

```rust
pub struct Schema {
    pub name: String,
    pub fields: Vec<Field>,
}

pub struct Field {
    pub name: String,
    pub data_type: String,
}
```

These can also handle nested data, but the json would look like this:

```json
{
    "name": "blog_post",
    "fields": [
        {
            "name": "title",
            "data_type": "string"
        },
        {
            "name": "body",
            "data_type": "string"
        },
        {
            "name": "date",
            "data_type": "date"
        }
    ]
}
```

#### Add data

1. select created schema
2. add data to the schema

#### Populdate UI with schema and data

1. get schema
2. create menu entry based on schema name and fill page with fields
3. get data based on schema name
4. fill list with data

This approach works for simple data types, but the following can't be handled:

- nested data
- relations
- files

Since files (images or videos) might be required for projects (or else), the Rust structs might be reworked to handle files. The files can be saved on a `Docker Volume` and the path can be saved in the database. But since all fields share the same types, a rework like this might be required:

```rust
pub struct Field {
    pub name: String,
    pub data_type: DataType,
}

pub enum DataType {
    String,
    Number,
    Boolean,
    Date,
    FileUrl,
}
```

The endpoints would then change from `web::json` to `Multipart`. If a schema contains the FileUrl data type, the endpoint needs to handle the file upload *(and writing to the volume)* first, then the data can be saved to the database *(with FileUrl updated with the path)*.

### WebUI

The implementation is not started, but might be implemented using NextJS or SvelteKit.
