use postgres::{Client, NoTls};

fn main() {
    // Read the second argument:
    let args: Vec<String> = std::env::args().collect();

    if args.len() != 2 {
        println!("Usage: {} <mode: loop>", args[0]);
        std::process::exit(1);
    }

    let mode = &args[1];

    match mode.as_ref() {
        "loop" => {
            let result = loop_();
            if let Err(e) = result {
                println!("Error running loop: {}", e);
                std::process::exit(1);
            }
        }
        _ => {
            println!("Unknown mode: {}", mode);
            std::process::exit(1);
        }
    }
}

fn loop_() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = Client::connect("host=localhost user=postgres", NoTls)?;

    for i in 0..200 {
        let mut transaction = client.transaction()?;

        let row = transaction.query_one("SELECT true", &[])?;
        let result: bool = row.get(0);

        println!("loop count {} result: {}", i, result);
    }

    Ok(())
}
