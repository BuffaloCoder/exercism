import re
from collections import Counter


def count_words(sentence):
    sentence = sentence.lower()
    words = re.findall(r"[a-z0-9]+(?:'\w)?", sentence)
    word_count = Counter(words)
    return word_count
