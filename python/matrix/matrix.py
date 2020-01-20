class Matrix(object):
    def __init__(self, matrix_string):
        self.matrix = []
        # split rows along line boundaries
        for row in matrix_string.splitlines():
            # generate the rows values  and add them to the list
            self.matrix.append([int(val) for val in row.split(' ')])

    def row(self, index):
        return self.matrix[index-1]

    def column(self, index):
        return [row[index-1] for row in self.matrix]
