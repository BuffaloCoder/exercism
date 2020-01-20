import re


def abbreviate(words):
    upper = str.upper(words)
    letters = [word[0] for word in re.split('[ _-]', upper) if len(word) > 0]
    return ''.join(letters)
