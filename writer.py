import sys
from datetime import datetime, timedelta
import subprocess
import os

a = [[0,0,1,1,1],
     [0,1,0,1,0],
     [1,0,0,1,0],
     [0,1,0,1,0],
     [0,0,1,1,1]]

b = [[1,1,1,1,1],
     [1,0,1,0,1],
     [1,0,1,0,1],
     [1,0,1,0,1],
     [0,1,0,1,0]]

c = [[0,1,1,1,0],
     [1,0,0,0,1],
     [1,0,0,0,1],
     [1,0,0,0,1],
     [1,0,0,0,1]]

d = [[1,1,1,1,1],
     [1,0,0,0,1],
     [1,0,0,0,1],
     [1,0,0,0,1],
     [0,1,1,1,0]]

e = [[1,1,1,1,1],
     [1,0,1,0,1],
     [1,0,1,0,1],
     [1,0,1,0,1],
     [1,0,0,0,1]]

f = [[1,1,1,1,1],
     [1,0,1,0,0],
     [1,0,1,0,0],
     [1,0,1,0,0],
     [1,0,0,0,0]]

g = [[0,1,1,1,0],
     [1,0,0,0,1],
     [1,0,1,0,1],
     [1,0,1,0,1],
     [1,0,1,1,0]]

h = [[1,1,1,1,1],
     [0,0,1,0,0],
     [0,0,1,0,0],
     [0,0,1,0,0],
     [1,1,1,1,1]]

i = [[1,0,0,0,1],
     [1,0,0,0,1],
     [1,1,1,1,1],
     [1,0,0,0,1],
     [1,0,0,0,1]]

j = [[1,0,0,1,0],
     [1,0,0,0,1],
     [1,0,0,0,1],
     [1,1,1,1,0],
     [1,0,0,0,0]]

k = [[1,1,1,1,1],
     [0,0,1,0,0],
     [0,0,1,0,0],
     [0,1,0,1,0],
     [1,0,0,0,1]]

l = [[1,1,1,1,1],
     [0,0,0,0,1],
     [0,0,0,0,1],
     [0,0,0,0,1],
     [0,0,0,0,1]]

m = [[1,1,1,1,1],
     [0,1,0,0,0],
     [0,0,1,0,0],
     [0,1,0,0,0],
     [1,1,1,1,1]]

n = [[1,1,1,1,1],
     [0,1,0,0,0],
     [0,0,1,0,0],
     [0,0,0,1,0],
     [1,1,1,1,1]]

o = [[0,1,1,1,0],
     [1,0,0,0,1],
     [1,0,0,0,1],
     [1,0,0,0,1],
     [0,1,1,1,0]]

p = [[1,1,1,1,1],
     [1,0,1,0,0],
     [1,0,1,0,0],
     [1,0,1,0,0],
     [0,1,0,0,0]]

q = [[0,1,1,1,0],
     [1,0,0,0,1],
     [1,0,1,0,1],
     [1,0,0,1,0],
     [0,1,1,0,1]]

r = [[1,1,1,1,1],
     [1,0,1,0,0],
     [1,0,1,0,0],
     [1,0,1,0,0],
     [0,1,0,1,1]]

s = [[0,1,0,0,1],
     [1,0,1,0,1],
     [1,0,1,0,1],
     [1,0,1,0,1],
     [1,0,0,1,0]]

t = [[1,0,0,0,0],
     [1,0,0,0,0],
     [1,1,1,1,1],
     [1,0,0,0,0],
     [1,0,0,0,0]]

u = [[1,1,1,1,0],
     [0,0,0,0,1],
     [0,0,0,0,1],
     [0,0,0,0,1],
     [1,1,1,1,0]]

v = [[1,1,1,0,0],
     [0,0,0,1,0],
     [0,0,0,0,1],
     [0,0,0,1,0],
     [1,1,1,0,0]]

w = [[1,1,1,1,1],
     [0,0,0,1,0],
     [0,0,1,0,0],
     [0,0,0,1,0],
     [1,1,1,1,1]]

x = [[1,0,0,0,1],
     [0,1,0,1,0],
     [0,0,1,0,0],
     [0,1,0,1,0],
     [1,0,0,0,1]]

y = [[1,0,0,0,0],
     [0,1,0,0,0],
     [0,0,1,1,1],
     [0,1,0,0,0],
     [1,0,0,0,0]]

z = [[1,0,0,0,1],
     [1,0,0,1,1],
     [1,0,1,0,1],
     [1,1,0,0,1],
     [1,0,0,0,1]]

alphabet = [a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z]

def make_git_commit(past_date):
    try:
        date_str = past_date.strftime('%Y-%m-%d')
        # Set the GIT_AUTHOR_DATE and GIT_COMMITTER_DATE environment variables
        env = {
            'GIT_AUTHOR_DATE': f'{date_str} 12:00:00',
            'GIT_COMMITTER_DATE': f'{date_str} 12:00:00',
            **dict(os.environ)
        }
        
        # Add all changes
        subprocess.run(['git', 'add', '.'], check=True)
        
        # Make the commit with the specified date
        subprocess.run(
            ['git', 'commit', '-m', f'contribution: {date_str}'],
            env=env,
            check=True
        )
    except subprocess.CalledProcessError as e:
        print(f"Error making git commit: {e}")
        sys.exit(1)

def print_letter(alphabet, index, past_date):
    for i in range(len(alphabet[index])):
        for j in range(len(alphabet[index][i])):
            if alphabet[index][i][j] == 1:
                past_date += timedelta(days=1)
                print(f"{past_date.strftime('%Y-%m-%d')}", end=" ")
                make_git_commit(past_date)
            else:
                past_date += timedelta(days=1)
                print(f"{past_date.strftime('%Y-%m-%d')}", end=" ")
                make_git_commit(past_date)
        print("character line")
    return past_date

def print_space(past_date):
    for i in range(7):
        past_date += timedelta(days=1)
        print(f"{past_date.strftime('%Y-%m-%d')}", end=" ")
        make_git_commit(past_date)
    print("space line")
    return past_date

def validate_input(text):
    if len(text) > 9:
        print("Error: Input must be 8 characters or less")
        sys.exit(1)
    
    if not text.isalpha():
        print("Error: Input must contain only letters a-z")
        sys.exit(1)
    
    if not text.islower():
        print("Error: Input must be lowercase letters only")
        sys.exit(1)

def print_word(word, past_date):
    # Convert letters to indices (a=0, b=1, etc)
    indices = [ord(c) - ord('a') for c in word]
    
    for i in range(len(indices)):
        past_date = print_letter(alphabet, indices[i], past_date)
        if i < len(indices) - 1:
            past_date = print_space(past_date)

if __name__ == "__main__":
    # Get date from 370 days ago
    today = datetime.now()
    past_date = today - timedelta(days=371)
    print(f"371 days ago was: {past_date.strftime('%Y-%m-%d')}")

    if len(sys.argv) != 2:
        print("Usage: python writer.py <word>")
        sys.exit(1)
        
    input_text = sys.argv[1].lower()
    validate_input(input_text)
    print_word(input_text, past_date)
