#!/usr/bin/python3

import random

num_of_test_inputs = [random.randint(1, 100) for i in range(0, 10)]
input_test_vectors = [[random.randint(-100, 100) for i in range(0, n)] for n in num_of_test_inputs]

def sum_of_squares_positive_values(test_vector):
    squares_of_positive = [i**2 for i in test_vector if i > 0]
    return sum(squares_of_positive)

result_vector = [sum_of_squares_positive_values(vector) for vector in input_test_vectors]

def encode_input_vector(input_vector):
    return ' '.join(map(str, input_vector))

print(result_vector)

# output to console, copy-paste to unit test file test param vector
[print(f'{{\"{encode_input_vector(k)}\", {v}}},') for k, v in zip(input_test_vectors, result_vector)]
