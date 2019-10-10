#[macro_use]
extern crate lazy_static;

mod socket;
mod cmd;

fn main() {
    socket::bootstrap_socket();
}
