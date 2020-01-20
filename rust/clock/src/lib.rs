use std::fmt;

#[derive(PartialEq, Debug)]
pub struct Clock {
    minutes: i32,
}

const HOURS_IN_A_DAY: i32 = 24;
const MINUTES_IN_AN_HOUR: i32 = 60;
const MINUTES_IN_A_DAY: i32 = MINUTES_IN_AN_HOUR * HOURS_IN_A_DAY;

fn calculate_time(current_minutes: i32, additional_minutes: i32) -> i32 {
    let mut minutes = (current_minutes + additional_minutes) % MINUTES_IN_A_DAY;
    if minutes < 0 {
        minutes += MINUTES_IN_A_DAY;
    }

    minutes
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let minutes = calculate_time(0, minutes + (hours * MINUTES_IN_AN_HOUR));
        Clock { minutes }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let minutes = calculate_time(self.minutes, minutes);
        Clock { minutes }
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(
            f,
            "{:02}:{:02}",
            self.minutes / MINUTES_IN_AN_HOUR,
            self.minutes % MINUTES_IN_AN_HOUR
        )
    }
}
