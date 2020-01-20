REPEATABLE_CHARS = {'-', ' '}

def is_isogram(string):
    chars_seen = set()
    lower_string = string.lower()
    for char in lower_string:
        # return early if the character is not repeatable and has
        # already been seen
        if char not in REPEATABLE_CHARS and char in chars_seen:
            return False
        chars_seen.add(char)
    return True
    
