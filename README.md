# <span style= "color:cyan">[ASCII-ART](https://www.zone01kisumu.ke/)</span>
Here is a masterpiece Generator, that takes text input and transforms it into stunning ASCII art using a series of intricate functions.
With the ASCII Art program, you can turn a simple text into a mesmerizing graphic masterpiece. 
 ## Files featured
The program works with 3 potent files:
 1. standard.txt
 2. shadow.txt
 3. thinkertoy.txt
 
 The default file used is standard textual file.

# <span style="color:cyan">How it Works</span> 
- The program takes a string input from the command line.
```bash
go run . hello
```
- It checks if the input string contains any unprintable characters (outside the ASCII range 32-126). If it does, it prints an error message and exits.eg.
```bash
go run . 你好
The program will exit in this casee
```
- It then checks if the input string contains any tabs (\t). If it does, it replaces them with four spaces ( ) eg.
```bash
go run . hello\\tworld
 _                                         
| |                                        
| |__                           __      __ 
|  _ \                          \ \ /\ / / 
| | | |                          \ V  V /  
|_| |_|                           \_/\_/   
```
- It then checks if the input string contains any newlines (\n). If it does,The program splits the input string into individual lines using the newline character (\n). eg.
```bash
go run . h\\nw
```
- For each line, it applies a set of predefined functions to convert each character to its corresponding ASCII art representation. The functions are:
    - FindHeadLine: maps characters to a specific ASCII art representation based on their position in the ASCII table.
    - Line2, Line3, Line4, Line5, Line6, Line7, and Line8: each maps characters to a specific ASCII art representation based on their position in the ASCII table.
- The program then concatenates the resulting ASCII art representations for each line to form the final output

## <span style="color:deeppink">Installation</span>

To install the program from Gitea, follow these steps:

- Clone the repository from Gitea using the provided URL in your terminal.
``` bash
$git clone https://learn.zone01kisumu.ke/git/rogwel/ascii-art
```
Ensure you have Go installed on your system. If not you could follow the link below:
```bash
https://go.dev/doc/install
```
Navigate to the project directory using the command:
    
```
$cd ascii-art
```

and run the program using:
```
go run . "YourTextHere"
```

# Files featured
The program works with 3 potent files:

 1. standard.txt
 2. shadow.txt
 3. thinkertoy.txt
 - The default file used is standard textual file.

 # <span style="color:cyan">The Final Stretch</span>

Congratulations, you've made it to the end of our README file! We hope you've found the information helpful in getting started with our project. As you embark on this exciting journey, remember to keep your wits about you, your code clean, and your coffee mug full. Don't hesitate to reach out if you have any questions or need further assistance.

Stay awesome, and happy coding!😊