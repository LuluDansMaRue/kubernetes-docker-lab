use serde::{Serialize, Deserialize};
use ws::{Message};
use crate::cmd::parser::{Kind, EnumHelper};

#[derive(Deserialize, Clone)]
pub struct Command {
    pub name: Option<String>,
    pub args: Option<Vec<String>>
}

#[derive(Deserialize, Clone)]
pub struct MessagePacket {
    pub id: String,
    pub sender_id: String,
    pub receiver_id: String,
    pub room_id: String,
    pub content: String,
    pub cmd: Command
}

#[derive(Serialize)]
pub struct ResponsePacket {
    pub room_id: String,
    pub content: String,
    pub sender_id: String
}

#[derive(Serialize)]
pub struct ResponseDataPacket {
    pub room_id: String,
    pub content_data: Vec<String>,
    pub sender_id: String
}

impl Command {
    pub fn retrieve_command(self) -> Kind {
        if let Some(name) = self.name {
            return Kind::from_string(&name);
        }

        Kind::Empty
    }
}

impl MessagePacket {
    pub fn create_from_message_packet(self) -> ResponsePacket {
        ResponsePacket {
            room_id: self.room_id,
            content: self.content,
            sender_id: self.sender_id
        }
    }
}

impl ResponsePacket {
    pub fn to_string(&self) -> String {
        match serde_json::to_string(&self) {
            Ok(json) => json,
            Err(err) => {
                // should log the error

                String::new()
            }
        }
    }
}

impl ResponseDataPacket {
    pub fn to_string(&self) -> String {
        match serde_json::to_string(&self) {
            Ok(json) => json,
            Err(err) => {
                // should log the error

                String::new()
            }
        }
    }
}

pub fn parse_message(msg: &Message) -> Option<MessagePacket> {
    if msg.is_empty() {
        return None;
    }

    if !msg.is_text() {
        // should log that other format are not support
        return None;
    }

    let content_str = match msg.as_text() {
        Ok(m) => m,
        Err(_) => {
            return None;
        }
    };

    println!("valeu of content {:?}", content_str);

    let packet: MessagePacket = match serde_json::from_str(&content_str) {
        Ok(json) => json,
        Err(err) => {
            // should log
            println!("error while parsing {:?}", err);
            return None;
        }
    };

    Some(packet)
}