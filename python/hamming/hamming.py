def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError('strands must be of the same length')

    return sum([1 for (a_gene, b_gene) in zip(strand_a, strand_b) if a_gene != b_gene])