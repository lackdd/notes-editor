Notes Tool
Overview
The Notes Tool is a simple command-line application designed to help you manage a collection of notes. You can view, add, or delete notes from a specified text file. If the file does not exist, the tool will create it automatically when you add a note. This tool is ideal for quickly managing text-based collections, such as coding ideas, task lists, or simple reminders.

Usage
To use the Notes Tool, you need to specify the name of the text file that will hold your notes when running the tool from the command line. Here’s how you can interact with the tool:

sh
Copy code
$ ./notestool <your_textfile_name>
Example
sh
Copy code
$ ./notestool coding_ideas
This command will load the tool, allowing you to manage notes stored in coding_ideas.txt. The tool will present a menu with the following options:

Show Notes: Display all the notes currently stored in the specified file.
Add a Note: Add a new note to the collection. If the file does not exist, it will be created.
Delete a Note: Delete a note from the collection by specifying its line number.
Exit: Exit the program.
Detailed Steps
Show Notes:

Select option 1 to view all notes in the collection. The notes are displayed line by line.
Add a Note:

Select option 2 to add a new note. You will be prompted to enter the note's text, which will then be appended to the file.
Delete a Note:

Select option 3 to delete a note. The tool will display all notes with corresponding numbers. Enter the number of the note you wish to delete.
Exit:

Select option 4 to exit the tool.
Data Storage
The notes are stored in a plain text file with the name you specify when running the tool. Each note is stored as a new line in this file. The tool appends new notes at the end of the file and can delete specific notes by removing the corresponding line.

Example of a Text File (coding_ideas.txt):
css
Copy code
Implement a binary search algorithm
Explore Go concurrency patterns
Build a REST API in Go
When you add a note, it’s added to the end of this list. When you delete a note, the tool rewrites the file without the selected line.