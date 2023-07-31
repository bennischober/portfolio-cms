use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, Deserialize, Serialize)]
pub struct User {
    pub first_name: String,
    pub last_name: String,
    pub username: String,
    pub email: String,
}

#[derive(Clone, Debug, PartialEq, Eq, Deserialize, Serialize)]
pub struct Schema {
    pub name: String,
    pub fields: Vec<Field>,
}

#[derive(Clone, Debug, PartialEq, Eq, Deserialize, Serialize)]
pub struct Field {
    pub name: String,
    pub data_type: DataType,
}

#[derive(Clone, Debug, PartialEq, Eq, Deserialize, Serialize)]
pub enum DataType {
    String,
    Number,
    Boolean,
    Date,
    FileUrl,
}
