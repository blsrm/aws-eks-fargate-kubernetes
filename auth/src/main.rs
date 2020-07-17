#[macro_use]
extern crate gotham_derive;
#[macro_use]
extern crate serde_derive;

use gotham::router::builder::*;
use gotham::router::Router;
use gotham::state::{FromState, State};


#[derive(Deserialize, StateData, StaticResponseExtender)]
struct PathExtractor {
  id: String,
}

fn current_user(state: State) -> (State, String) {
  let message = {
    let currentuser = PathExtractor::borrow_from(&state);
    format!("currentuser: {}", currentuser.id)
  };
  (state, message)
}

const PONG: &str = "Pong!";
fn ping(state: State) -> (State, &'static str) {
  (state, PONG)
}

const SIGNUP: &str = "user signed up";
fn sign_up(state: State) -> (State, &'static str) {
  (state, SIGNUP)
}

const SIGNIN: &str = "user signed in";
fn sign_in(state: State) -> (State, &'static str) {
  (state, SIGNIN)
}

const SIGNOUT: &str = "user signed out";
fn sign_out(state: State) -> (State, &'static str) {
  (state, SIGNOUT)
}


fn router() -> Router {
  build_simple_router(|route| {
    route.get("/ping").to(ping);

    route.get("/currentuser/:id")
    .with_path_extractor::<PathExtractor>()
    .to(current_user);

    route.post("/signup").to(sign_up);
    route.post("/signin").to(sign_in);
    route.post("/signout").to(sign_out);
  })
}

pub fn main() {
  let addr = "127.0.0.1:7878";
  println!("Listening for requests at http://{}", addr);
  gotham::start(addr, router())
}

