# PROJECT NAME : ASCII-ART

# Description:


This program is a Go application that converts a string into ASCII art. It supports various characters, including letters, numbers, spaces, and special characters within the printable range. You can also customize the output style using different banner files.

# Features:



    Input String: Accepts a string as input from the command line.
    Character Handling: Supports a wide range of printable characters (letters, numbers, spaces, and special characters).
    Newline Handling: Preserves newlines (\n) within the input string for multi-line ASCII art.
    Error Handling: Detects non-printable characters and issues encountered with the banner file.
    Banner Customization: Allows using different banner files (e.g., standard.txt, thinkertoy.txt) to change the ASCII art style.




# Usage:


    Running the Program: Open a terminal in your project directory and execute the following command:
    Bash

    go run . "your_string_here" [banner_file.txt]

    Use code with caution.

        Replace your_string_here with the text you want to convert to ASCII art.
        Optional: Specify a banner file (e.g., standard.txt, thinkertoy.txt) to customize the output style. Defaults to standard.txt if omitted.

Examples:
Bash

Print "Hello World" with the standard banner
go run . "Hello World" | cat -e
 _    _          _   _                __          __                 _       _  $
| |  | |        | | | |               \ \        / /                | |     | | $
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | $
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | $
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | $
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| $
                                                                                $
                                                                                $
> 

Print "Hello" with the thinkertoy banner
go run . "Hello" thinkertoy.txt | cat -e
                 $
o  o     o o     $
|  |     | |     $
O--O o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $

#Print a multi-line message with the shadow banner (assuming a shadow.txt exists)
go run . "Hello\nThere" shadow.txt | cat -e
                                $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
                                               $
_|_|_|_|_| _|                                  $
    _|     _|_|_|     _|_|   _|  _|_|   _|_|   $
    _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
    _|     _|    _| _|       _|       _|       $
    _|     _|    _|   _|_|_| _|         _|_|_| $
                                               $
                                               $

Use code with caution.



