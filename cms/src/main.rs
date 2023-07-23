mod model;
use model::{User, Schema};

use actix_web::{get, post, web, App, HttpResponse, HttpServer};
use bson::{doc, Bson, Document};
use mongodb::{Client, Collection};
//use serde::Serialize;
//use serde_json::Value;

const DB_NAME: &str = "myApp";
const COLL_NAME: &str = "users";


/// Gets the user with the supplied username.
#[get("/api/get_user/{username}")]
async fn get_user(client: web::Data<Client>, username: web::Path<String>) -> HttpResponse {
    let username = username.into_inner();
    let collection: Collection<User> = client.database(DB_NAME).collection(COLL_NAME);
    match collection
        .find_one(doc! { "username": &username }, None)
        .await
    {
        Ok(Some(user)) => HttpResponse::Ok().json(user),
        Ok(None) => {
            HttpResponse::NotFound().body(format!("No user found with username {username}"))
        }
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

/// Creates a new schema.
#[post("/api/schema")]
async fn create_schema(client: web::Data<Client>, schema: web::Json<Schema>) -> HttpResponse {
    // Convert the Value to a Document
    let doc = match bson::to_bson(&schema.into_inner()) {
        Ok(Bson::Document(document)) => document,
        Ok(_) => return HttpResponse::BadRequest().json("Invalid data"),
        Err(_) => return HttpResponse::InternalServerError().json("Internal server error"),
    };

    // Insert the Document into MongoDB
    let result = insert_into_mongodb(client.get_ref(), doc).await;
    match result {
        Ok(()) => HttpResponse::Ok().finish(),
        Err(_) => HttpResponse::InternalServerError().json("Failed to insert into MongoDB"),
    }
}

/// Gets the schema by name
#[get("/api/schema/{name}")]
async fn get_schema(client: web::Data<Client>, name: web::Path<String>) -> HttpResponse {
    let name = name.into_inner();
    let collection: Collection<Document> = client.database("test_db").collection("test_col");
    match collection.find_one(doc! { "name": &name }, None).await {
        Ok(Some(schema)) => {
            match serde_json::to_string(&schema) {
                Ok(json) => HttpResponse::Ok().body(json),
                Err(_) => HttpResponse::InternalServerError().body("Error converting document to JSON"),
            }
        },
        Ok(None) => {
            HttpResponse::NotFound().body(format!("No schema found with name {}", name))
        }
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

/// Inserts a Document into MongoDB.
async fn insert_into_mongodb(client: &Client, doc: Document) -> mongodb::error::Result<()> {
    // Access a specific collection (change "my_db" and "my_collection" to your actual DB and collection names)
    let collection = client.database("test_db").collection("test_col");

    // Insert our document into the collection
    collection.insert_one(doc, None).await?;

    Ok(())
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let uri = std::env::var("MONGODB_URI").unwrap_or_else(|_| "mongodb://localhost:27017".into());
    let client = Client::with_uri_str(uri).await.expect("failed to connect");

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(client.clone()))
            .service(get_user)
            .service(create_schema)
            .service(get_schema)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
