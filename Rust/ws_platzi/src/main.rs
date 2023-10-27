#[macro_use]
extern crate diesel;

pub mod schema;
pub mod models;

use dotenv::dotenv;
use std::env;
use diesel::prelude::*;
use diesel::pg::PgConnection;
use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder, Error};

#[get("/hello")]
async fn hello_world() -> impl Responder {
    HttpResponse::Ok().body("holo")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    HttpServer::new(|| {
        App::new().service(hello_world)
    }).bind(("0.0.0.0", 9900))?.run().await
    // let db_url = env::var("DATABASE_URL").expect("db url does not found");

    // let conn = PgConnection::establish(&db_url).expect("No se pudo conectar a la base de datos");

    // use self::models::{Post, NewPost, PostSimplificado};
    // use self::schema::posts;
    // use self::schema::posts::dsl::*;

    // let new_post = NewPost{
    //     title: "Mi tercer blogsot",
    //     body: "coso",
    //     slug: "primer-post",
    // };

    // let _post: Post = diesel::insert_into(posts::table).values(&new_post).get_result(&conn).expect("La insertada fallo");

    // update
    // let _post_update = diesel::update(posts.filter(id.eq(3))).set(slug.eq("tercer-post")).get_result::<Post>(&conn).expect("error en el update");

    // delete
    // diesel::delete(posts.filter(id.eq(3))).execute(&conn).expect("ha fallado la eliminacion del tercer post");

    // SELECT * FROM posts
    // let posrt_results = posts.load::<Post>(&conn).expect("Error al ejecutar la query");

    // for post in posrt_results {
    //     println!("{:?}", post);
    // }

    // // solo algunios
    // println!("query con columnas especificas");
    // let posts_results = posts.select((title, body)).limit(1).load::<PostSimplificado>(&conn).expect("Error al ejecutar la query");

    // for post in posts_results {
    //     println!("{:?}", post);
    // }
}
