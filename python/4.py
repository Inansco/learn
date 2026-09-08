word = input("Enter a word: ")

print("===== STRING ANALYZER =====\n")

print(f"Word: {word}")
print(f"First character: {word[0]}")
print(f"Last character: {word[-1]}")
print(f"First 3 characters: {word[:3]}")
print(f"Last 3 characters: {word[-3:]}")
print(f"Reversed: {word[::-1]}")

print(f"Length: {len(word)}")