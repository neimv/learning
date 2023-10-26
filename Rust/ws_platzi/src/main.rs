#[macro_use]
extern crate diesel;

pub mod schema;
pub mod models;

use dotenv::dotenv;
use std::env;
use diesel::prelude::*;
use diesel::pg::PgConnection;

fn main() {
    dotenv().ok();
    let db_url = env::var("DATABASE_URL").expect("db url does not found");

    let conn = PgConnection::establish(&db_url).expect("No se pudo conectar a la base de datos");

    use self::models::{Post, NewPost};
    use self::schema::posts;
    use self::schema::posts::dsl::*;

    let new_post = NewPost{
        title: "Mi primer blogsot",
        body: "coso",
        slug: "primer-post",
    };

    let _post: Post = diesel::insert_into(posts::table).values(&new_post).get_result(&conn).expect("La insertada fallo");

    // SELECT * FROM posts
    let posrt_results = posts.load::<Post>(&conn).expect("Error al ejecutar la query");

    for post in posrt_results {
        println!("{}", post.title);
    }
}
