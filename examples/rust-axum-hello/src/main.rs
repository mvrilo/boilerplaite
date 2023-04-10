use axum::{handler::get, Router};
use clap::{App, Arg};

async fn hello_world() -> &'static str {
    "Hello, World!"
}

#[tokio::main]
async fn main() {
    let matches = App::new("Hello World")
        .version("0.1.0")
        .author("Your Name <you@example.com>")
        .about("Prints 'Hello, world!'")
        .arg(
            Arg::new("port")
                .short('p')
                .about("Sets the port")
                .takes_value(true),
        )
        .get_matches();

    let port: u16 = matches.value_of("port").unwrap_or("3000").parse().unwrap();

    let app = Router::new().route("/", get(hello_world));

    axum::Server::bind(&format!("0.0.0.0:{}", port).parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}