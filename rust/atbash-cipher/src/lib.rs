/// "Encipher" with the Atbash cipher.
pub fn encode(plain: &str) -> String {
    let mut encryption = String::from("");
    // for each char
    for c in plain.chars() {
        // if it's non-alphanumeric, throw it away
        if !c.is_alphanumeric() {
            continue;
        }
        // if the length+1 is divisible by 6, we need a space
        if ((encryption.len() + 1) % 6) == 0 {
            encryption.push(' ')
        }
        // if it's numeric, keep it
        if c.is_numeric() {
            encryption.push(c)
        } else {
            // else, swap it out with it's reverse
            let reverse_bytes = b'z' - c.to_ascii_lowercase() as u8 + b'a';
            encryption.push(reverse_bytes as char);
        }
    }
    encryption
}

/// "Decipher" with the Atbash cipher.
pub fn decode(cipher: &str) -> String {
    let mut decryption = String::from("");
    // for each char
    for c in cipher.chars() {
        // if it's a space, skip it
        if c == ' ' {
            continue;
        }

        // if it's numeric, keep it
        if c.is_numeric() {
            decryption.push(c)
        } else {
            // else, swap it out with it's reverse
            let reverse_bytes = b'z' - c.to_ascii_lowercase() as u8 + b'a';
            decryption.push(reverse_bytes as char);
        }
    }
    decryption
}
