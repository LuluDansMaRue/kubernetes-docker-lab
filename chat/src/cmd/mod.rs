pub mod parser {
    /// Trait to parse to get enum from string
    pub trait EnumHelper<T, E> {
        fn from_string(action: &String) -> E;
    }

    pub enum Kind {
        Room,
        Bobba,
        Insult,
        Empty
    }

    impl EnumHelper<String, Kind> for Kind {
        fn from_string(kind: &String) -> Kind {      
            match kind.to_lowercase().as_str() {
                "room" => Kind::Room,
                "bobba" => Kind::Bobba,
                "insult" => Kind::Insult,
                _ => Kind::Empty
            }
        }
    }

    pub enum Actions {
        Create,
        Join,
        List,
        Delete,
        Order,
        Empty
    }

    impl EnumHelper<String, Actions> for Actions {
        fn from_string(action: &String) -> Actions {      
            match action.to_lowercase().as_str() {
                "create" => Actions::Create,
                "join" => Actions::Join,
                "list" => Actions::List,
                "delete" => Actions::Delete,
                "order" => Actions::Order,
                _ => Actions::List
            }
        }
    }

    /// Get Action
    /// 
    /// # Description
    /// Parse the arguments of a command
    /// e.g: /room create <name>
    /// 
    /// # Arguments
    /// * `args` String
    pub fn get_action(args: Option<Vec<String>>) -> (Actions, Vec<String>) {
        if let None = args {
            return (Actions::Empty, Vec::new());
        }
        
        let mut opts = args.unwrap();
        let action_str = match opts.first() {
            Some(a) => a,
            None => {
                return (Actions::Empty, Vec::new());
            }
        };

        let action = Actions::from_string(action_str);
        opts.remove(0);

        (action, opts)
    }
}