mod room;
mod message;

use ws::{listen, CloseCode, Handler, Handshake, Message, Result, Sender};
use crate::cmd::parser::{Kind, get_action};

#[derive(Clone)]
struct Server {
    out: Sender,
}

impl Handler for Server {
    fn on_open(&mut self, _: Handshake) -> Result<()> {
        // Join default channel
        room::join_room(String::from("default"), self.out.clone());
        room::message_default_room(None);

        Ok(())
    }

    fn on_message(&mut self, msg: Message) -> Result<()> {
        if msg.is_empty() {
            return Ok(());
        }

        let msg_packet = message::parse_message(&msg);
        if let Some(msg) = msg_packet {
            println!("parse packet");
            // we should avoid use clone but for the sake of an example...
            let kind = msg.clone().cmd.retrieve_command();
            let (action, args) = get_action(msg.clone().cmd.args);

            match kind {
                Kind::Bobba => {},
                Kind::Room => room::handle_room_actions_from_cmd(action, args, self.out.clone()),
                Kind::Insult => {},
                Kind::Empty => room::broadacast_message_to_room(msg)
            }
        }

        Ok(())
    }

    fn on_close(&mut self, code: CloseCode, reason: &str) {
        match code {
            CloseCode::Normal => println!("The client is done with the connection."),
            CloseCode::Away   => println!("The client is leaving the site."),
            _ => println!("The client encountered an error: {}", reason),
        }
    }
}

pub fn bootstrap_socket() {
    room::create_default_room();
    println!("Socket server running on port 8088");

    listen("127.0.0.1:8088", |out| {
        Server {
            out: out
        }
    } ).unwrap();
}