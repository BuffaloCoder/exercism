FIRST_VERSE = 'a Partridge in a Pear Tree.'
GIFTS = [
    'two Turtle Doves',
    'three French Hens',
    'four Calling Birds',
    'five Gold Rings',
    'six Geese-a-Laying',
    'seven Swans-a-Swimming',
    'eight Maids-a-Milking',
    'nine Ladies Dancing',
    'ten Lords-a-Leaping',
    'eleven Pipers Piping',
    'twelve Drummers Drumming'
]

DAYS = [
    'first',
    'second',
    'third',
    'fourth',
    'fifth',
    'sixth',
    'seventh',
    'eighth',
    'ninth',
    'tenth',
    'eleventh',
    'twelfth'
]

def recite(start_verse: int, end_verse: int):
    output = []
    for verse_index in range(start_verse, end_verse+1):
        verse = "On the {} day of Christmas my true love gave to me: ".format(DAYS[verse_index-1])
        verse += ', '.join([GIFTS[x] for x in reversed(range(0, verse_index-1))])
        verse += ', and ' if (verse_index > 1) else ''
        verse += FIRST_VERSE
        output.append(verse)
    return output
