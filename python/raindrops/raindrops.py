RAIN_DESCRIPTORS = [
    {
        'check': 3,
        'result': 'Pling'
    },
    {
        'check': 5,
        'result': 'Plang'
    },
    {
        'check': 7,
        'result': 'Plong'
    }
]

def raindrops(number):
    drops = [desc['result'] for desc in RAIN_DESCRIPTORS if number % desc['check'] == 0]
    return ''.join(drops) if drops else str(number)