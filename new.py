text_input = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
max_width = 42
words = text_input.split()
lines = []
current_line = ""

for word in words:
    if len(current_line) + len(word) + 1 > max_width:
        lines.append(current_line)
        current_line = word
    else:
        if current_line:
            current_line += " " + word
        else:
            current_line = word

if current_line:
    lines.append(current_line)

for line in lines:
    print(line)
