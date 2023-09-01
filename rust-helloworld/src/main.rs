use axum::{
    routing::get,
    Router, http::Method,
};
use tower_http::cors::{CorsLayer, Any};

#[tokio::main]
async fn main() {
    let cors = CorsLayer::new()
        .allow_methods([Method::GET])
        .allow_origin(Any);

    // build our application with a single route
    let app = Router::new()
        .layer(cors)
        .route("/", get(|| async { "Hello, 世界!" }));

    // run it with hyper on localhost:3000
    axum::Server::bind(&"0.0.0.0:3000".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}