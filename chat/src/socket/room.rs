use std::collections::HashMap;
use std::sync::Mutex;
use std::cell::RefCell;
use serde_json;
use ws::{Sender, Message};
use crate::socket::message::{MessagePacket, ResponsePacket, ResponseDataPacket};
use crate::cmd::parser::Actions;

#[derive(Debug)]
pub struct Room {
    pub room_name: String,
    pub senders: RefCell<Vec<Sender>>
}

lazy_static! {
    #[derive(Debug)]
    pub static ref ROOMS: Mutex<HashMap<String, Room>> = Mutex::new(HashMap::new());
}

impl Room {
    pub fn send_to_room(&mut self, msg: Message) {
        let senders: &Vec<Sender> = self.senders.get_mut();
        for sender in senders {
            let res = sender.send(msg.clone());
            match res {
                Ok(_) => (),
                Err(_) => ()
            }
        }
    }

    pub fn send_welcome_payload(&mut self) {
        let packet = ResponsePacket {
            room_id: String::from("default"),
            content: String::from("Welcome to the bobba chat"),
            sender_id: String::from("all")
        };

        let msg = Message::text(packet.to_string());
        self.send_to_room(msg);
    }
}

pub fn create_default_room() {
    ROOMS.lock().unwrap().insert(
        String::from("default"),
        Room {
            room_name: String::from("default"),
            senders: RefCell::new(Vec::new())
        }
    );
}

/// DIRTY AF
pub fn join_room(room_id: String, sender: Sender) {
    let mut exist = false;
    let welcome_packet = ResponsePacket {
        room_id: String::from(&room_id),
        content: format!("Welcome to the channel: {}", room_id),
        sender_id: String::from("all")
    };

    let room_opt = ROOMS.lock();
    if let Ok(r) = room_opt {
        let channel_opt = r.get(&room_id);
        if let Some(channel) = channel_opt {
            let packet = Message::text(welcome_packet.to_string());
            let res = sender.send(packet);
            if let Err(err) = res {
                // should print the error
                return;
            }

            let senders: &mut Vec<Sender> = &mut channel.senders.borrow_mut();
            for s in senders.iter() {
                if s.connection_id() == sender.connection_id() {
                    exist = true;
                }
            }

            if exist {
                return;
            }

            match res {
                Ok(_) => senders.push(sender),
                Err(err) => panic!(err)
            };
        }
    }
}

pub fn broadacast_message_to_room(packet: MessagePacket) {
    let room_id = packet.clone().room_id;

    let rooms_mutex = ROOMS.lock();
    if let Err(err) = rooms_mutex {
        // log something
        return;
    }

    let mut rooms = rooms_mutex.unwrap();
    let room = rooms.get_mut(&room_id);

    if let Some(r) = room {
        let resp_packet = packet.create_from_message_packet();
        let msg = Message::text(resp_packet.to_string());
        r.send_to_room(msg);
    }
}

pub fn message_default_room(payload: Option<String>) {
    let mut lock = ROOMS.lock().unwrap();
    let room = lock.get_mut(&String::from("default"));
    // nested if are bad but anyway for demo
    if let Some(r) = room {
        if let Some(p) = payload {
            let msg = Message::text(p);
            r.send_to_room(msg);

            return;
        }

        r.send_welcome_payload();
    }
}

pub fn create_room(name: Option<&String>) {
    println!("creating room");
    let room_name = match name {
        Some(n) => String::from(n),
        None => String::new()
    };

    ROOMS.lock().unwrap().insert(
        room_name.clone(),
        Room {
            room_name: room_name,
            senders: RefCell::new(Vec::new())
        }
    );
}

pub fn get_room_names() -> Option<String> {
    let hashmap = match ROOMS.lock() {
        Ok(l) => l,
        Err(err) => {
            // should log
            return None;
        }
    };

    let rooms_name: Vec<String> = hashmap
        .iter()
        .map(|(key, _)| String::from(key))
        .collect();
    
    // rooms_name
    let packet = ResponseDataPacket {
        room_id: String::from("default"),
        content_data: rooms_name,
        sender_id: String::from("bobba")
    };

    let json = match serde_json::to_string(&packet) {
        Ok(j) => j,
        Err(err) => {
            // should log the error
            println!("error while creating json room list {:?}", err);
            return None;
        }
    };

    Some(json)
}

pub fn handle_room_actions_from_cmd(action: Actions, arg: Vec<String>, sender: Sender) {
    println!("handling room action");

    match action {
        Actions::Create => create_room(arg.get(0)),
        Actions::Join => {
            match arg.get(0) {
                Some(id_string) => {
                    join_room(String::from(id_string), sender);
                },
                None => {}
            }
        },
        Actions::List => {
            let list = get_room_names();
            message_default_room(list);
        },
        _ => println!("nada")
    }
}