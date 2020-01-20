pub fn is_armstrong_number(num: u32) -> bool {
    // Make a copy of the number to digest
    let mut num_copy = num;

    // Get the digits
    let mut digits = Vec::new();
    let mut digit_count = 0;
    while num_copy > 0 {
        digits.push(num_copy % 10);
        num_copy /= 10;
        digit_count += 1;
    }

    // Calculate the armstrong number
    let mut armstrong_number = 0;
    for digit in &digits {
        armstrong_number += digit.pow(digit_count);
    }

    // Compare the armstrong number to the original
    armstrong_number == num
}
