import os


directory = os.getcwd() + "/testdata/bench"

file_name = "large_file.txt"

total_lines = 1000000


os.makedirs(directory, exist_ok=True)


file_path = os.path.join(directory, file_name)


with open(file_path, 'w', encoding='utf-8') as f:
    for i in range(1, total_lines):
        f.write(f"This is line {i}\n")